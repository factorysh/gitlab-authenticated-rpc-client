package main

import (
	"os"

	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/cli"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/command"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/display"
)

func main() {
	app := cli.NewApp()
	cmd := cli.NewClient()

	cli.Register(app,
		command.NewAuthClient(cmd),
		command.NewGitlabClient(cmd),
	)

	err := app.Run(os.Args)
	if err != nil {
		display.Error("Error %v", err)
	}

}
