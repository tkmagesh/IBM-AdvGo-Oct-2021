package main

import (
	"context"
	"grpc-app/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	response := &proto.AddResponse{
		Sum: result,
	}
	return response, nil
}

func (s *server) Average(stream proto.AppService_AverageServer) error {
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

func main() {

	//Hosting the service
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grcpServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grcpServer, &server{})
	grcpServer.Serve(listener)
}
