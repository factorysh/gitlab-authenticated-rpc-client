package main

import (
	"os"

	"github.com/factorysh/gitlab-authenticated-rpc-client/cli"
	"github.com/factorysh/gitlab-authenticated-rpc-client/command"
	"github.com/factorysh/gitlab-authenticated-rpc-client/display"
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
