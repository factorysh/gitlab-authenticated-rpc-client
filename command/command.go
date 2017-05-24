package command

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	Conn *grpc.ClientConn
	Ctx  context.Context
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		Conn: conn,
		Ctx:  context.Background(),
	}
}
