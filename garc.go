package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"runtime"

	"golang.org/x/net/context"

	"gitlab.bearstech.com/bearstech/journaleux/gar/client/auth"
	"gitlab.bearstech.com/bearstech/journaleux/gar/client/conf"
	"gitlab.bearstech.com/bearstech/journaleux/gar/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":50051"
)

var (
	git_version = ""
)

func main() {
	domain := os.Args[1]

	cfg := conf.NewConf("gar", domain)
	t, err := cfg.Token()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(t)
	// Set up a connection to the server.
	conn, err := grpc.Dial(domain+port,
		grpc.WithPerRPCCredentials(&auth.IdAuth{Token: t}),
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
	md := metadata.Pairs()

	hello, err := h.SayHello(ctx, &rpc.HelloRequest{os.Args[2]}, grpc.Trailer(&md))
	if err != nil {
		log.Fatalf("Can't hello: %v %v\n", err, md)
	}
	log.Println(hello)
	hello, err = h.SayHello(ctx, &rpc.HelloRequest{"Super " + os.Args[2]}, grpc.Trailer(&md))
	if err != nil {
		log.Fatalf("Can't hello: %v %v\n", err, md)
	}
	log.Println(hello)

}
