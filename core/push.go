package core

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/segmentio/ksuid"
	"tawny/tawny"
)

func Subscribe(subscribeInput *tawny.SubscribeInput, stream tawny.PushService_SubscribeServer) error {
	id, err := ksuid.NewRandom()
	fmt.Printf("subscribe channel:%s topic:%s id:%s \n", subscribeInput.Channel, subscribeInput.Topic, id.String())
	if err != nil {
		fmt.Println(err)
		return err
	}
	c := make(chan tawny.Message)
	if val, ok := SubscriptionStore[subscribeInput.Topic+subscribeInput.Channel]; ok {
		val[id.String()] = &c
	} else {
		SubscriptionStore[subscribeInput.Topic+subscribeInput.Channel] = map[string]*chan (tawny.Message){
			id.String(): &c,
		}
	}
	defer func() {
		delete(SubscriptionStore[subscribeInput.Topic+subscribeInput.Channel], id.String())
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
			fmt.Printf("subscribe channel:%s topic:%s id:%s message \n", subscribeInput.Channel, subscribeInput.Topic, id.String())
		}
	}
	return nil
}
func Publish(ctx context.Context, pushInput *tawny.PushInput) (em *empty.Empty, err error) {
	em = &empty.Empty{}
	if val, ok := SubscriptionStore[pushInput.Topic+pushInput.Channel]; ok {
		for _, e := range val {
			m := tawny.Message{
				Data: pushInput.Data,
			}
			fmt.Printf("push channel:%s topic:%s message \n", pushInput.Channel, pushInput.Topic)
			*e <- m
		}
	}
	return
}
