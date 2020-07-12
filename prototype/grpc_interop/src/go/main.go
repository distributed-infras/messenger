package main

import (
	pb "../grpc/message/my.org/message"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	// "net/http"
	// "io/ioutil"
	// "encoding/json"
)

func main() {
	grpcServer := grpc.NewServer()
	var server Server
	pb.RegisterMessageServer(grpcServer, server)
	listen, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:3000 %v", err)
	}
	log.Println("Server starting...")
	log.Fatal(grpcServer.Serve(listen))
}

type Server struct{}

func (Server) Search(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{Payload: "Hello"}, nil
}
