package command

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/urfave/cli"
)

func (c *GitlabClient) User(_cli *cli.Context) error {
	err := c.Client.SetDomain(_cli.GlobalString("domain"))
	if err != nil {
		return err
	}
	g := c.rpcClient()
	_, err = g.Ping(c.Client.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	u, err := g.MyUser(c.Client.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	fmt.Printf("User:\n\t%+v\n", u)
	return nil
}

func registerUsers(g *GitlabClient, app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "Get yourself",
		Action:  g.User,
	})
}
