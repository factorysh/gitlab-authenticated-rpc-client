package main

import (
	"crypto/tls"
	"log"
	"os"

	"golang.org/x/net/context"

	"gitlab.bearstech.com/bearstech/journaleux/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

func main() {
	domain := os.Args[1]

	// Set up a connection to the server.
	conn, err := grpc.Dial(domain+port,
		//grpc.WithPerRPCCredentials(&DummyAuth{}),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	h := rpc.NewHelloServiceClient(conn)

	ctx := context.Background()

	hello, err := h.SayHello(ctx, &rpc.HelloRequest{os.Args[2]})
	if err != nil {
		log.Fatalf("Can't hello: %v", err)
	}
	log.Println(hello)

}
