package cli

import (
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	Conn *grpc.ClientConn
	Ctx  context.Context
}

func NewClient() *Client {
	return &Client{
		Ctx: context.Background(),
	}
}

func (c *Client) SetDomain(domain string) (err error) {
	cc := client.New(domain)
	c.Conn, err = cc.ClientConn(true)
	return err
}
