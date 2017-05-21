package main

import (
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/skratchdot/open-golang/open"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/client"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	md := metadata.Pairs()

	hello, err := h.SayHello(ctx, &rpc.HelloRequest{os.Args[2]}, grpc.Trailer(&md))
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Println(st)
		}
		u, ok := md["gar.auth_code_url"]
		if ok {
			if !strings.HasPrefix(u[0], "https://") {
				log.Fatal("Bad url prefix, we all gonna die")
			}
			open.Run(u[0])
		}
		log.Fatalf("Can't hello: %v %v\n", err, md)
	}
	log.Printf("Hello: %s\n", hello)
	hello, err = h.SayHello(ctx, &rpc.HelloRequest{"Super " + os.Args[2]}, grpc.Trailer(&md))
	if err != nil {
		log.Fatalf("Can't hello: %v %v\n", err, md)
	}
	log.Printf("Super Hello: %s\n", hello)
	u, err := h.WhoAmI(ctx, &empty.Empty{}, grpc.Trailer(&md))
	if err != nil {
		log.Fatalf("Can't who am I: %v %v\n", err, md)
	}
	log.Printf("Who am I?: %v\n", u)

}
