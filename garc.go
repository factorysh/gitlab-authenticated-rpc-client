package main

import (
	"log"
	"os"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/client"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
)

const (
	port = ":50051"
)

var (
	git_version = ""
)

func main() {
	domain := os.Args[1]

	conn, err := client.NewConn(domain)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	h := rpc.NewHelloServiceClient(conn)

	ctx := context.Background()

	hello, err := h.SayHello(ctx, &rpc.HelloRequest{os.Args[2]})
	if err != nil {
		log.Fatalf("Can't hello: %v\n", err)
	}
	log.Printf("Hello: %s\n", hello)
	hello, err = h.SayHello(ctx, &rpc.HelloRequest{"Super " + os.Args[2]})
	if err != nil {
		log.Fatalf("Can't hello: %v\n", err)
	}
	log.Printf("Super Hello: %s\n", hello)
	u, err := h.WhoAmI(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Can't who am I: %v \n", err)
	}
	log.Printf("Who am I?: %v\n", u)

}
