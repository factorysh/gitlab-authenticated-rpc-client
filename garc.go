package main

import (
	"os"

	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/command"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/display"
)

func main() {
	app := command.NewApp()

	cmd := command.NewClient()
	aa := command.NewAuthClient(cmd)
	gg := command.NewGitlabClient(cmd)

	aa.Register(app)
	gg.Register(app)
	err := app.Run(os.Args)
	if err != nil {
		display.Error("Error %v", err)
	}

}
