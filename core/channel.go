package core

import (
	"context"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
	"tawny/tawny"
)

func VerifyChannel(name string, ctx context.Context) (configuration *tawny.ChannelConfiguration, err error) {
	configuration = &tawny.ChannelConfiguration{}
	Db.View(func(txn *badger.Txn) error {
		key := getAdminKey(name)

		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			err = proto.Unmarshal(val, configuration)
			return err
		})
		return nil
	})

	//md, ok := metadata.FromIncomingContext(ctx)
	//if !ok {
	//	fmt.Println("NOT OK METADATA")
	//}
	//keys := md.Get("API_KEY")
	//if len(keys) != 1 {
	//
	//}
	//key := keys[0]

	return
}

func IsAdmin(ctx context.Context) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("NOT OK METADATA")
	}
	keys := md.Get("API_KEY")
	if len(keys) != 1 {
		return false
	}
	key := keys[0]
	return key == viper.GetString(ADMIN_KEY)
}
