package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

//var reqCount int = 0

type server struct {
	sync.Mutex
	reqCount int64
	proto.UnimplementedAppServiceServer
}

func (s *server) IncrementCount() {
	//atomic.AddInt64(&s.reqCount, 1)
	s.Lock()
	{
		s.reqCount++
	}
	s.Unlock()
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	//reqCount++
	//s.reqCount++
	s.IncrementCount()
	x := req.GetX()
	y := req.GetY()
	result := x + y
	response := &proto.AddResponse{
		Sum: result,
	}
	return response, nil
}

func (s *server) Average(stream proto.AppService_AverageServer) error {
	//s.reqCount++
	s.IncrementCount()
	var sum int32
	var count int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//all the requests are received
			//prepare the response and send it (use SendAndClose())
			avg := sum / count
			response := &proto.AverageResponse{
				Result: avg,
			}
			log.Println("Average Request : sending response result = ", avg)
			return stream.SendAndClose(response)
		}
		if err != nil {
			return err
		}
		no := req.GetNo()
		log.Println("Average Request : No = ", no)
		sum = sum + no
		count++
	}
	return nil
}

func (s *server) GeneratePrime(req *proto.PrimeRequest, stream proto.AppService_GeneratePrimeServer) error {
	//s.reqCount++
	s.IncrementCount()
	start, end := req.GetStart(), req.GetEnd()
	log.Println("GeneratePrime Request : No = ", start, end)
	for i := start; i <= end; i++ {
		if isPrime(i) {
			time.Sleep(500 * time.Millisecond)
			log.Println("Sending Prime : ", i)
			response := &proto.PrimeResponse{
				No: i,
			}
			stream.Send(response)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	if no <= 1 {
		return false
	}
	var i int32
	for i = 2; i < no; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (s *server) GreetEveryone(stream proto.AppService_GreetEveryoneServer) error {

	for {
		//s.reqCount++

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		s.IncrementCount()
		firstName := req.GetUser().GetFirstName()
		lastName := req.GetUser().GetLastName()
		message := fmt.Sprintf("Hi %s %s!", firstName, lastName)
		time.Sleep(1 * time.Second)
		log.Println("Sending Greeting : ", message)
		response := &proto.GreetResponse{
			Message: message,
		}
		stream.Send(response)
	}
}

func main() {
	s := &server{}
	go func() {
		for {
			s.Lock()
			{
				log.Println("Request Count : ", s.reqCount)
			}
			s.Unlock()
			//log.Println("Request Count : ", atomic.LoadInt64(&s.reqCount))
			time.Sleep(1 * time.Minute)
		}
	}()

	//Hosting the service
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grcpServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grcpServer, s)
	grcpServer.Serve(listener)
}
