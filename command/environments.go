package command

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/urfave/cli"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
)

func (c *GitlabClient) Environments(_cli *cli.Context) error {
	err := c.Client.SetDomain(_cli.GlobalString("domain"))
	if err != nil {
		return err
	}
	var pid string
	if _cli.Args().First() != "" {
		pid = _cli.Args().First()
	}
	if pid == "" {
		return errors.New("Invalid project name")
	}
	g := c.rpcClient()
	_, err = g.Ping(c.Client.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	envs, err := g.MyEnvironments(c.Client.Ctx, &rpc.ProjectPredicate{Id: pid})
	if err != nil {
		return err
	}
	fmt.Printf("Environments:\n")
	for _, p := range envs.Environments {
		fmt.Printf("\t%+v\n", p)
	}
	return nil
}

func registerEnvironements(g *GitlabClient, app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    "environments",
		Aliases: []string{"e"},
		Usage:   "Get your environments for a project",
		Action:  g.Environments,
	})
}
