package main

import (
	"github.com/urfave/cli"
	"os"
	"sort"

	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/command"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/display"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "Gitlab authenticated rpc client"
	app.Version = version.GitVersion
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "domain, d",
			Value: "rpc.example.com",
			Usage: "Target RPC server",
		},
	}

	cmd := command.NewClient()
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

	sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		display.Error("Error %v", err)
	}

}
