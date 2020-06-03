package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"strings"
	"tawny/tawny"
)

type AdminServer struct {
}

const adminPrefix = "admin@@@"

func getAdminKey(channelName string) []byte {
	return []byte(fmt.Sprintf("%s%s", adminPrefix, channelName))
}

func (s *AdminServer) CreateChannelOrUpdate(ctx context.Context, input *tawny.CreateOrUpdateChannelInput) (e *empty.Empty, err error) {
	err = Db.Update(func(txn *badger.Txn) error {
		key := getAdminKey(input.Name)
		var b []byte
		b, err = proto.Marshal(input.Configuration)
		e := badger.NewEntry(key, b)
		err := txn.SetEntry(e)
		return err
	})
	e = &empty.Empty{}
	return
}
func (s *AdminServer) ListChannel(ctx context.Context, empty *empty.Empty) (output *tawny.ListChannelOutput, err error) {
	channels := []*tawny.Channel{}
	err = Db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(adminPrefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			keyRaw := item.Key()
			name := strings.TrimPrefix(string(keyRaw), adminPrefix)
			err := item.Value(func(v []byte) error {
				var config tawny.ChannelConfiguration
				err := proto.Unmarshal(v, &config)
				if err != nil {
					return err
				}
				channels = append(channels, &tawny.Channel{
					Name:          name,
					Configuration: &config,
				})
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	output = &tawny.ListChannelOutput{
		Channels: channels,
	}
	return
}

func (s *AdminServer) DeleteChannel(ctx context.Context, input *tawny.DeleteChannelInput) (e *empty.Empty, err error) {
	var channelConfig *tawny.ChannelConfiguration
	channelConfig, err = VerifyChannel(input.Name, ctx)
	if channelConfig != nil {
		err = errors.New(fmt.Sprintf("channel %s do not exist or inaccessible", input.Name))
	}
	err = Db.Update(func(txn *badger.Txn) error {
		key := getAdminKey(input.Name)
		err := txn.Delete(key)
		return err
	})
	e = &empty.Empty{}
	return
}
