package main

import (
	"context"
	"log"
	"net"

	"github.com/GoSG/livedemo/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GreetMe(ctx context.Context, req *proto.GreetRequest) (resp *proto.GreetResponse, err error) {
	var response string

	firstName := req.GetFirstName()
	lastName := req.GetLastName()

	response = "Hello " + firstName + " " + lastName

	resp = &proto.GreetResponse{
		Response: response,
	}
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal("error occured")
	}

	s := grpc.NewServer()
	proto.RegisterGreetServer(s, &server{})

	if err = s.Serve(listen); err != nil {
		log.Fatal("error occured while")
	}
}
