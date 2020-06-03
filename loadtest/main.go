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

	conn, err := grpc.Dial("tawny.bobby-demo.site:4000", grpc.WithInsecure())
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

	}()
	end := make(chan bool)

	for i := 0; i < 50; i++ {
		go Publish(strconv.Itoa(i))
		go Publish(strconv.Itoa(i))

		go Subscriber(strconv.Itoa(i), end)
		go Subscriber(strconv.Itoa(i), end)

	}
	for i := 0; i < 49; i++ {
		<-end
	}
	defer conn.Close()
}

func Subscriber(topic string, end chan bool) {
	md := metadata.New(map[string]string{"API_KEY": "abc"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	conn, err := grpc.Dial("tawny.bobby-demo.site:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	pushClient := tawny.NewPushServiceClient(conn)
	psb, err := pushClient.Subscribe(ctx, &tawny.SubscribeInput{
		Channel: "test-load",
		Topic:   topic,
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
		if message == 200 {
			psb.CloseSend()
			break
		}
		if message%10 == 0 {
			fmt.Printf("Result: call %d latency_average: %d max:%d \n", message, latency/message, max)
		}
	}
	fmt.Printf("Result: call %d latency_average: %d max:%d \n", message, latency/message, max)
	conn.Close()
	end <- true
}

func Publish(topic string) {
	md := metadata.New(map[string]string{"API_KEY": "abc"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	conn, err := grpc.Dial("tawny.bobby-demo.site:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	pushClient := tawny.NewPushServiceClient(conn)
	for i := 0; i < 200; i++ {
		time.Sleep(time.Millisecond * 50)
		_, err := pushClient.Publish(ctx, &tawny.PushInput{
			Channel: "test-load",
			Topic:   topic,
			Data:    []byte(strconv.FormatInt(time.Now().UnixNano(), 10)),
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
