package command

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/urfave/cli"
)

func (c *AuthClient) Ping(_cli *cli.Context) error {
	err := c.Client.SetDomain(_cli.GlobalString("domain"))
	if err != nil {
		return err
	}
	a := c.rpcClient()
	_, err = a.Ping(c.Client.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	fmt.Println("Pong")
	return nil
}

func registerPing(a *AuthClient, app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:   "ping",
		Usage:  "Ping the server, without auth",
		Action: a.Ping,
	})
}
