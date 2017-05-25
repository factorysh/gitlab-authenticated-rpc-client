package command

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
	c.Conn, err = client.NewConn(domain)
	return err
}
