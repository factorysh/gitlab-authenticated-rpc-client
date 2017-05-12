package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"runtime"

	"golang.org/x/net/context"

	"gitlab.bearstech.com/bearstech/journaleux/gar/client/auth"
	"gitlab.bearstech.com/bearstech/journaleux/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

var (
	git_version = ""
)

func main() {
	domain := os.Args[1]

	// Set up a connection to the server.
	conn, err := grpc.Dial(domain+port,
		grpc.WithPerRPCCredentials(&auth.JWTAuth{Token: "password"}),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
		grpc.WithUserAgent(fmt.Sprintf("Journaleux %s #%s", runtime.GOOS, git_version)),
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
