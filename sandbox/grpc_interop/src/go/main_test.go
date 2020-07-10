package main

import (
	pb "../grpc/message/my.org/message"
	"golang.org/x/net/context"
	// "google.golang.org/grpc"
	"log"
	"testing"
	// "net"
	// "net/http"
	// "io/ioutil"
	// "encoding/json"
)

func TestMessage(t *testing.T) {
	ctx := context.Background()
	req := pb.Request{Payload: ""}
	server := Server{}
	res, err := server.Search(ctx, &req)
	if err != nil {
		t.Error(err)
	}
	if res.Payload != "Hello" {
		t.Error("Response is incorrect")
	}
	log.Println(res)
}
