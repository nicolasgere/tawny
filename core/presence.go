package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/segmentio/ksuid"
	"strings"
	"sync"
	"tawny/tawny"
	"time"
)

var mutexPresence = sync.Mutex{}

type PresenceServer struct {
}

func (s *PresenceServer) HeartBeat(ctx context.Context, heartBeatInput *tawny.HeartBeatInput) (em *empty.Empty, err error) {
	return HeartBeat(ctx, heartBeatInput)
}
func (s *PresenceServer) PresenceSubscribe(presenceInput *tawny.PresenceSubscribeInput, stream tawny.PresenceService_PresenceSubscribeServer) error {
	return PresenceSubscribe(presenceInput, stream)
}

func getPresence(input *tawny.PresenceSubscribeInput) (data []*tawny.Presence, err error) {
	key := input.Topic + input.Channel
	data = []*tawny.Presence{}
	err = Db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(key)

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			keyRaw := item.Key()
			key := strings.Split(string(keyRaw), "%%%%")[1]
			err := item.Value(func(v []byte) error {
				copy := append([]byte{}, v...)
				data = append(data, &tawny.Presence{
					State: copy,
					Key:   key,
				})
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

func PresenceSubscribe(input *tawny.PresenceSubscribeInput, stream tawny.PresenceService_PresenceSubscribeServer) (err error) {
	var channelConfig *tawny.ChannelConfiguration
	channelConfig, err = VerifyChannel(input.Channel, stream.Context())
	if channelConfig != nil {
		err = errors.New(fmt.Sprintf("channel %s do not exist or inaccessible", input.Channel))
	}

	id, err := ksuid.NewRandom()
	key := input.Topic + input.Channel
	fmt.Printf("presence channel:%s topic:%s id:%s \n", input.Channel, input.Topic, id.String())
	if err != nil {
		fmt.Println(err)
		return err
	}
	c := make(chan bool)
	mutexPresence.Lock()
	if val, ok := PresenceStore[key]; ok {
		val[id.String()] = &c
	} else {
		PresenceStore[key] = map[string]*chan (bool){
			id.String(): &c,
		}
	}
	mutexPresence.Unlock()

	defer func() {
		mutexPresence.Lock()
		delete(PresenceStore[key], id.String())
		mutexPresence.Unlock()
	}()
	ticker := time.Tick(time.Second * 10)
Loop:
	for {
		select {
		case <-c:
			data, err := getPresence(input)
			if err != nil {
				fmt.Println(err)
				break Loop
			}
			err = stream.Send(&tawny.PresenceSubscribeResponse{
				Presences:  data,
				UpdateType: tawny.PresenceSubscribeResponse_FULL})
			if err != nil {
				fmt.Println(err)
				break Loop
			}
			fmt.Printf("presence channel:%s topic:%s id:%s message \n", input.Channel, input.Topic, id.String())
		case <-ticker:
			data, err := getPresence(input)
			if err != nil {
				fmt.Println(err)
				break Loop
			}
			err = stream.Send(&tawny.PresenceSubscribeResponse{
				Presences:  data,
				UpdateType: tawny.PresenceSubscribeResponse_FULL,
			})
			if err != nil {
				fmt.Println(err)
				break Loop
			}
			fmt.Printf("presence channel:%s topic:%s id:%s message \n", input.Channel, input.Topic, id.String())
		}
	}
	return nil
}
func HeartBeat(ctx context.Context, heartBeatInput *tawny.HeartBeatInput) (em *empty.Empty, err error) {
	em = &empty.Empty{}
	key := heartBeatInput.Topic + heartBeatInput.Channel
	err = Db.Update(func(txn *badger.Txn) error {

		e := badger.NewEntry([]byte(key+"%%%%"+heartBeatInput.Key), heartBeatInput.State).WithTTL(time.Second * 10)
		err := txn.SetEntry(e)
		return err
	})

	for _, e := range PresenceStore[key] {
		*e <- true
	}
	return
}
