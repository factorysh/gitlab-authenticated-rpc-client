package command

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/urfave/cli"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc_auth"
)

func (c *Client) Ping(_cli *cli.Context) error {
	err := c.SetDomain(_cli.GlobalString("domain"))
	if err != nil {
		return err
	}
	a := auth.NewAuthClient(c.Conn)
	_, err = a.Ping(c.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	return nil
}
