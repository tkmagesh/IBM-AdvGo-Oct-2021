package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	/* doRequestResponse(ctx, client)
	doClientStreaming(ctx, client)
	doServerStreaming(ctx, client)
	doBidiStreaming(ctx, client) */
	doRequestResponseWithTimeout(ctx, client)
}

func doRequestResponseWithTimeout(ctx context.Context, client proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	timeoutCtx, cancelFn := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancelFn()
	resp, err := client.Add(timeoutCtx, addRequest)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Timeout error")
			} else {
				log.Fatalln(err)
			}
		}
		log.Fatalln(err)
	}
	log.Println("Result : ", resp.GetSum())
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

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	stream, err := client.GeneratePrime(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("Received all responses")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Prime : ", res.GetNo())
	}
}

func doBidiStreaming(ctx context.Context, client proto.AppServiceClient) {
	users := []proto.User{
		{FirstName: "Magesh", LastName: "Kuppan"},
		{FirstName: "Suresh", LastName: "Rajan"},
		{FirstName: "Rajesh", LastName: "Pandit"},
		{FirstName: "Ramesh", LastName: "Jayaraman"},
		{FirstName: "Ganesh", LastName: "Kumar"},
	}
	stream, err := client.GreetEveryone(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		for _, user := range users {
			req := &proto.GreetRequest{
				User: &user,
			}

			time.Sleep(5 * time.Second)
			log.Println("Sending : ", fmt.Sprintf("%v", user))
			stream.Send(req)
		}
		log.Println("Sent all the requests")
	}()
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("Received all responses")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Message : ", res.GetMessage())
		}
		done <- true
	}()
	<-done
}
