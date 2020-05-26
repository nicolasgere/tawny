package core

import (
	"context"
	"github.com/dgraph-io/badger/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"tawny/tawny"
)

var SubscriptionStore = map[string]map[string]*chan (tawny.Message){}
var PresenceStore = map[string]map[string]*chan (bool){}

var db *badger.DB

func Init() {
	var err error
	db, err = badger.Open(badger.DefaultOptions("db-temp/badger"))
	if err != nil {
		log.Fatal(err)
	}
}

type PushServer struct {
}

func (s *PushServer) Subscribe(subscribeInput *tawny.SubscribeInput, stream tawny.PushService_SubscribeServer) error {
	return Subscribe(subscribeInput, stream)
}
func (s *PushServer) Publish(ctx context.Context, pushInput *tawny.PushInput) (*empty.Empty, error) {
	return Publish(ctx, pushInput)
}

type PresenceServer struct {
}

func (s *PresenceServer) HeartBeat(ctx context.Context, heartBeatInput *tawny.HeartBeatInput) (em *empty.Empty, err error) {
	return HeartBeat(ctx, heartBeatInput)
}
func (s *PresenceServer) PresenceSubscribe(presenceInput *tawny.PresenceSubscribeInput, stream tawny.PresenceService_PresenceSubscribeServer) error {
	return PresenceSubscribe(presenceInput, stream)
}
