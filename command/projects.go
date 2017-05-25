package command

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/urfave/cli"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
)

func (c *Client) Projects(_cli *cli.Context) error {
	err := c.SetDomain(_cli.GlobalString("domain"))
	if err != nil {
		return err
	}
	g := rpc.NewGitlabClient(c.Conn)
	_, err = g.Ping(c.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	pp, err := g.MyProjects(c.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	fmt.Printf("Projects:\n")
	for _, p := range pp.Projects {
		fmt.Printf("\t%+v\n", p)
	}
	return nil
}
