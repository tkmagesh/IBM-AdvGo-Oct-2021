package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	//doRequestResponse(ctx, client)
	doClientStreaming(ctx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}

	resp, err := client.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Result : ", resp.GetSum())
}

func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		data := []int32{6, 3, 7, 5, 2, 4, 8, 1, 9}
		clientStream, err := client.Average(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		//sending multiple requests to the server
		for _, val := range data {
			averageRequest := &proto.AverageRequest{
				No: val,
			}
			log.Println("	Sending (1): No = ", val)
			time.Sleep(500 * time.Millisecond)
			clientStream.Send(averageRequest)
		}

		//receiving the response from the server
		res, err := clientStream.CloseAndRecv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("	Average (1): ", res.GetResult())
		wg.Done()
	}()

	go func() {
		data := []int32{60, 30, 70, 50, 20, 40, 80, 10, 90}
		clientStream, err := client.Average(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		//sending multiple requests to the server
		for _, val := range data {
			averageRequest := &proto.AverageRequest{
				No: val,
			}
			log.Println("Sending (2): No = ", val)
			time.Sleep(250 * time.Millisecond)
			clientStream.Send(averageRequest)
		}

		//receiving the response from the server
		res, err := clientStream.CloseAndRecv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Average (2): ", res.GetResult())
		wg.Done()
	}()
	wg.Wait()
}
