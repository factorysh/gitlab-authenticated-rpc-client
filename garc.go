package main

import (
	"github.com/urfave/cli"
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
	var domain string
	ctx := context.Background()

	app := cli.NewApp()
	app.Name = "Gitlab authenticated rpc client"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "domain, d",
			Value:       "rpc.example.com",
			Usage:       "Target RPC server",
			Destination: &domain,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Get yourself",
			Action: func(c *cli.Context) error {
				conn, err := client.NewConn(domain)
				if err != nil {
					log.Fatal(err)
				}
				g := rpc.NewGitlabClient(conn)
				_, err = g.Ping(ctx, &empty.Empty{})
				if err != nil {
					return err
				}
				u, err := g.MyUser(ctx, &empty.Empty{})
				if err != nil {
					return err
				}
				log.Println("User: ", u)
				return nil
			},
		},
		{
			Name:    "projects",
			Aliases: []string{"p"},
			Usage:   "Get your projects",
			Action: func(c *cli.Context) error {
				conn, err := client.NewConn(domain)
				if err != nil {
					log.Fatal(err)
				}
				defer conn.Close()
				g := rpc.NewGitlabClient(conn)
				_, err = g.Ping(ctx, &empty.Empty{})
				if err != nil {
					return err
				}
				pp, err := g.MyProjects(ctx, &empty.Empty{})
				if err != nil {
					return err
				}
				for _, p := range pp.Projects {
					log.Println("\tProject: ", p)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)

}
