package command

import (
	_cli "github.com/factorysh/gitlab-authenticated-rpc-client/cli"
	"github.com/factorysh/gitlab-authenticated-rpc-client/rpc_gitlab"
	"github.com/urfave/cli"
)

// GitlabClient is client for Gitlab rpc
type GitlabClient struct {
	Client *_cli.Client
}

// NewGitlabClient return a new GitlabClient
func NewGitlabClient(client *_cli.Client) *GitlabClient {
	return &GitlabClient{Client: client}
}

func (g *GitlabClient) rpcClient() rpc.GitlabClient {
	return rpc.NewGitlabClient(g.Client.Conn)
}

func (g *GitlabClient) Register(app *cli.App) {
	registerEnvironements(g, app)
	registerProjects(g, app)
	registerUsers(g, app)
}
