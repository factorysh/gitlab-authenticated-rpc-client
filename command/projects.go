package command

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/urfave/cli"
)

func (c *GitlabClient) Projects(_cli *cli.Context) error {
	err := c.Client.SetDomain(_cli.GlobalString("domain"))
	if err != nil {
		return err
	}
	g := c.rpcClient()
	_, err = g.Ping(c.Client.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	pp, err := g.MyProjects(c.Client.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	fmt.Printf("Projects:\n")
	for _, p := range pp.Projects {
		fmt.Printf("\t%+v\n", p)
	}
	return nil
}

func registerProjects(g *GitlabClient, app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    "projects",
		Aliases: []string{"p"},
		Usage:   "Get your projects",
		Action:  g.Projects,
	})
}
