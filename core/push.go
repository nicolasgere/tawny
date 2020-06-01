package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/segmentio/ksuid"
	"sync"
	"tawny/tawny"
)

var mutex = sync.Mutex{}

type PushServer struct {
}

func (s *PushServer) Subscribe(subscribeInput *tawny.SubscribeInput, stream tawny.PushService_SubscribeServer) error {
	return Subscribe(subscribeInput, stream)
}
func (s *PushServer) Publish(ctx context.Context, pushInput *tawny.PushInput) (*empty.Empty, error) {
	return Publish(ctx, pushInput)
}

func Subscribe(input *tawny.SubscribeInput, stream tawny.PushService_SubscribeServer) (err error) {
	var channelConfig *tawny.ChannelConfiguration
	channelConfig, err = VerifyChannel(input.Channel, stream.Context())
	if channelConfig != nil {
		err = errors.New(fmt.Sprintf("channel %s do not exist or inaccessible", input.Channel))
	}
	id, err := ksuid.NewRandom()
	fmt.Printf("subscribe channel:%s topic:%s id:%s \n", input.Channel, input.Topic, id.String())
	if err != nil {
		fmt.Println(err)
		return err
	}
	c := make(chan tawny.Message)
	mutex.Lock()
	if val, ok := SubscriptionStore[input.Topic+input.Channel]; ok {
		val[id.String()] = &c
	} else {
		SubscriptionStore[input.Topic+input.Channel] = map[string]*chan (tawny.Message){
			id.String(): &c,
		}
	}
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(SubscriptionStore[input.Topic+input.Channel], id.String())
		mutex.Unlock()

	}()
Loop:
	for {
		select {
		case m := <-c:
			err := stream.Send(&m)
			if err != nil {
				fmt.Println(err)
				break Loop
			}
			fmt.Printf("subscribe channel:%s topic:%s id:%s message \n", input.Channel, input.Topic, id.String())
		}
	}
	return nil
}
func Publish(ctx context.Context, input *tawny.PushInput) (em *empty.Empty, err error) {
	var channelConfig *tawny.ChannelConfiguration
	channelConfig, err = VerifyChannel(input.Channel, ctx)
	if channelConfig == nil {
		err = errors.New(fmt.Sprintf("channel %s do not exist or inaccessible", input.Channel))
	}
	if err != nil {
		return
	}
	isAdmin := IsAdmin(ctx)
	if channelConfig.AdminRequiredToPush && isAdmin == false {
		err = errors.New(fmt.Sprintf("channel %s require admin credential to push", input.Channel))
	}
	em = &empty.Empty{}
	mutex.Lock()
	val, ok := SubscriptionStore[input.Topic+input.Channel]
	mutex.Unlock()
	if ok {
		for _, e := range val {
			m := tawny.Message{
				Data: input.Data,
			}
			fmt.Printf("push channel:%s topic:%s message \n", input.Channel, input.Topic)
			*e <- m
		}
	}
	return
}
