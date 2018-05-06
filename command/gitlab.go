package command

import (
	"github.com/urfave/cli"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
)

// GitlabClient is client for Gitlab rpc
type GitlabClient struct {
	Client *Client
}

// NewGitlabClient return a new GitlabClient
func NewGitlabClient(client *Client) *GitlabClient {
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
