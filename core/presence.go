package core

import (
	"context"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/segmentio/ksuid"
	"strings"
	"tawny/tawny"
	"time"
)

func getPresence(presenceSubscribeInput *tawny.PresenceSubscribeInput) (data []*tawny.Presence, err error) {
	key := presenceSubscribeInput.Topic + presenceSubscribeInput.Channel
	data = []*tawny.Presence{}
	err = db.View(func(txn *badger.Txn) error {
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

func PresenceSubscribe(presenceInput *tawny.PresenceSubscribeInput, stream tawny.PresenceService_PresenceSubscribeServer) error {
	id, err := ksuid.NewRandom()
	key := presenceInput.Topic + presenceInput.Channel

	fmt.Printf("presence channel:%s topic:%s id:%s \n", presenceInput.Channel, presenceInput.Topic, id.String())
	if err != nil {
		fmt.Println(err)
		return err
	}
	c := make(chan bool)
	if val, ok := PresenceStore[key]; ok {
		val[id.String()] = &c
	} else {
		PresenceStore[key] = map[string]*chan (bool){
			id.String(): &c,
		}
	}
	defer func() {
		delete(PresenceStore[key], id.String())
	}()
	ticker := time.Tick(time.Second * 10)
Loop:
	for {
		select {
		case <-c:
			data, err := getPresence(presenceInput)
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
			fmt.Printf("presence channel:%s topic:%s id:%s message \n", presenceInput.Channel, presenceInput.Topic, id.String())
		case <-ticker:
			data, err := getPresence(presenceInput)
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
			fmt.Printf("presence channel:%s topic:%s id:%s message \n", presenceInput.Channel, presenceInput.Topic, id.String())
		}
	}
	return nil
}
func HeartBeat(ctx context.Context, heartBeatInput *tawny.HeartBeatInput) (em *empty.Empty, err error) {
	em = &empty.Empty{}
	key := heartBeatInput.Topic + heartBeatInput.Channel
	err = db.Update(func(txn *badger.Txn) error {

		e := badger.NewEntry([]byte(key+"%%%%"+heartBeatInput.Key), heartBeatInput.State).WithTTL(time.Second * 10)
		err := txn.SetEntry(e)
		return err
	})

	for _, e := range PresenceStore[key] {
		*e <- true
	}
	return
}
