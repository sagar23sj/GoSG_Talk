package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GoSG/livedemo/proto"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("0.0.0.0:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("error occured while creating connection")
	}

	defer cc.Close()

	conn := proto.NewGreetClient(cc)

	GreetMe(conn)
}

func GreetMe(c proto.GreetClient) {
	req := proto.GreetRequest{
		FirstName: "sagar",
		LastName:  "sonwane",
	}

	rsp, err := c.GreetMe(context.Background(), &req)
	if err != nil {
		log.Fatal("error calling Greet method")
	}

	fmt.Println("Response From server : ", rsp.Response)
}
