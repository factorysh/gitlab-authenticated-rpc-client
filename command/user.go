package command

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/urfave/cli"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
)

func (c *Client) User(_cli *cli.Context) error {
	err := c.SetDomain(_cli.GlobalString("domain"))
	if err != nil {
		return err
	}
	g := rpc.NewGitlabClient(c.Conn)
	_, err = g.Ping(c.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	u, err := g.MyUser(c.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	fmt.Printf("User:\n\t%+v\n", u)
	return nil
}
