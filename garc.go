package main

import (
	"os"
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

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
			Name:   "domain, d",
			Value:  "rpc.example.com:50051",
			Usage:  "Target RPC server",
			EnvVar: "DOMAIN",
		},
		cli.BoolFlag{
			Name:  "verbose, vv",
			Usage: "Log verbosity",
		},
	}

	cmd := command.NewClient()
	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("verbose") {
			log.SetLevel(log.DebugLevel)
			log.Debug("Verbose")
		} else {
			log.SetLevel(log.InfoLevel)
		}
		return nil
	}
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
		{
			Name:    "environments",
			Aliases: []string{"e"},
			Usage:   "Get your environments for a project",
			Action:  cmd.Environments,
		},
		{
			Name:   "ping",
			Usage:  "Ping the server, without auth",
			Action: cmd.Ping,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		display.Error("Error %v", err)
	}

}
