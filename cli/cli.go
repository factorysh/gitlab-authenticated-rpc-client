package cli

import (
	"sort"

	log "github.com/sirupsen/logrus"
	_cli "github.com/urfave/cli"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/version"
)

func NewApp() *_cli.App {
	app := _cli.NewApp()
	app.Name = "Gitlab authenticated rpc client"
	app.Version = version.GitVersion
	app.EnableBashCompletion = true

	app.Flags = []_cli.Flag{
		_cli.StringFlag{
			Name:   "domain, d",
			Value:  "rpc.example.com:50051",
			Usage:  "Target RPC server",
			EnvVar: "DOMAIN",
		},
		_cli.BoolFlag{
			Name:  "verbose, vv",
			Usage: "Log verbosity",
		},
	}

	app.Before = func(c *_cli.Context) error {
		if c.GlobalBool("verbose") {
			log.SetLevel(log.DebugLevel)
			log.Debug("Verbose")
		} else {
			log.SetLevel(log.InfoLevel)
		}
		return nil
	}

	sort.Sort(_cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))
	return app
}

// Registrable CLI
type Registrable interface {
	Register(app *_cli.App)
}

func Register(app *_cli.App, registrables ...Registrable) {
	for _, registrable := range registrables {
		registrable.Register(app)
	}
}
