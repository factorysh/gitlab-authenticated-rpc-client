package command

import (
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/client"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
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
	ca, err := conf.GetCA(os.Getenv("CA_PATH"))
	if err != nil {
		return err
	}
	cc := client.New(domain)
	cc.CertPool = ca
	c.Conn, err = cc.ClientConn()
	return err
}
