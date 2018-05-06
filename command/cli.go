package command

import (
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/version"
)

func NewApp() *cli.App {
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

	cmd := NewClient()
	aa := NewAuthClient(cmd)
	gg := NewGitlabClient(cmd)

	RegisterAuth(aa, app)
	RegisterGitlab(gg, app)

	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("verbose") {
			log.SetLevel(log.DebugLevel)
			log.Debug("Verbose")
		} else {
			log.SetLevel(log.InfoLevel)
		}
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))
	return app
}
