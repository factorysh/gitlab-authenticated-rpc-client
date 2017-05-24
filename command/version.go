package command

import (
	"fmt"
	"github.com/urfave/cli"
)

var (
	GitVersion = "🦇 "
)

func (c *Client) Version(_cli *cli.Context) error {
	fmt.Println(GitVersion)
	return nil
}
