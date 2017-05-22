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
	g := rpc.NewGitlabClient(conn)

	ctx := context.Background()
	_, err = g.Ping(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Can't ping: %v\n", err)
	}

	u, err := g.MyUser(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Can't get my user: %v\n", err)
	}
	log.Println("User: ", u)

}
