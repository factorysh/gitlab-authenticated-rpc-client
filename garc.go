package main

import (
	"github.com/urfave/cli"
	"log"
	"os"

	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/client"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/command"
)

const (
	port = ":50051"
)

var (
	git_version = ""
)

func main() {
	var domain string

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

	conn, err := client.NewConn(domain)
	if err != nil {
		log.Fatal(err)
	}
	cmd := command.NewClient(conn)
	app.Commands = []cli.Command{
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Get yourself",
			Action:  cmd.User,
		},
		{
			Name:    "projects",
			Aliases: []string{"p"},
			Usage:   "Get your projects",
			Action:  cmd.Projects,
		},
	}

	app.Run(os.Args)

}
