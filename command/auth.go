package command

import (
	"github.com/urfave/cli"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc_auth"
)

type AuthClient struct {
	Client *Client
}

func NewAuthClient(client *Client) *AuthClient {
	return &AuthClient{Client: client}
}

func (a *AuthClient) rpcClient() auth.AuthClient {
	return auth.NewAuthClient(a.Client.Conn)
}

func (a *AuthClient) Register(app *cli.App) {
	registerPing(a, app)
}
