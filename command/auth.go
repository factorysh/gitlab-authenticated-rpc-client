package command

import (
	_cli "github.com/factorysh/gitlab-authenticated-rpc-client/cli"
	"github.com/factorysh/gitlab-authenticated-rpc-client/rpc_auth"
	"github.com/urfave/cli"
)

type AuthClient struct {
	Client *_cli.Client
}

func NewAuthClient(client *_cli.Client) *AuthClient {
	return &AuthClient{Client: client}
}

func (a *AuthClient) rpcClient() auth.AuthClient {
	return auth.NewAuthClient(a.Client.Conn)
}

func (a *AuthClient) Register(app *cli.App) {
	registerPing(a, app)
}
