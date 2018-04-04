package client

import (
	"context"
	"os"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	c := New(os.Getenv("DOMAIN"))
	g, err := c.NewGitlabClient()
	assert.Nil(t, err)
	ctx := context.Background()
	r, err := g.Ping(ctx, &empty.Empty{})
	assert.Nil(t, err)
	assert.Equal(t, &empty.Empty{}, r)
}
func TestMyUser(t *testing.T) {
	c := New(os.Getenv("DOMAIN"))
	g, err := c.NewGitlabClient()
	assert.Nil(t, err)
	ctx := context.Background()
	u, err := g.MyUser(ctx, &empty.Empty{})
	assert.Nil(t, err)
	t.Log(u)
}
