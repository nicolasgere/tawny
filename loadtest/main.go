package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
	"tawny/tawny"
	"time"
)

func main() {

	md := metadata.New(map[string]string{"API_KEY": "abc"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	adminClient := tawny.NewAdminServiceClient(conn)

	adminClient.CreateChannelOrUpdate(ctx, &tawny.CreateOrUpdateChannelInput{
		Name: "test-load",
		Configuration: &tawny.ChannelConfiguration{
			AdminRequiredToPush: false,
		},
	})
	go func() {
		pushClient := tawny.NewPushServiceClient(conn)
		for i := 0; i < 2000; i++ {
			time.Sleep(1 * time.Millisecond)
			_, err := pushClient.Publish(ctx, &tawny.PushInput{
				Channel: "test-load",
				Topic:   "1",
				Data:    []byte(strconv.FormatInt(time.Now().UnixNano(), 10)),
			})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
	end := make(chan bool)
	go Subscriber(ctx, conn, end)
	go Subscriber(ctx, conn, end)
	go Subscriber(ctx, conn, end)
	go Subscriber(ctx, conn, end)
	<-end
	<-end
	<-end
	<-end
	defer conn.Close()
}

func Subscriber(ctx context.Context, conn *grpc.ClientConn, end chan bool) {

	pushClient := tawny.NewPushServiceClient(conn)
	psb, err := pushClient.Subscribe(ctx, &tawny.SubscribeInput{
		Channel: "test-load",
		Topic:   "1",
	})
	if err != nil {
		panic(err)
	}
	message := int64(0)
	latency := int64(0)
	max := int64(0)
	for {
		msg, err := psb.Recv()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		now := time.Now().UnixNano()
		sendAt, err := strconv.ParseInt(string(msg.Data), 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		message++
		diff := (now - sendAt) / 1e6
		latency = latency + diff
		if diff > max {
			max = diff
		}
		if message == 2000 {
			break
		}
	}
	fmt.Printf("Result: call %d latency_average: %d max:%d \n", message, latency/message, max)
	end <- true
}
