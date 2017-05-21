package client

import (
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/auth"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"runtime"
)

const (
	port = ":50051"
)

var (
	git_version = ""
)

func NewConn(domain string) (*grpc.ClientConn, error) {

	cfg := conf.NewConf("gar", domain)
	t, err := cfg.Token()
	if err != nil {
		return nil, errors.Wrap(err, "Can't get token")
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(domain+port,
		grpc.WithPerRPCCredentials(&auth.IdAuth{Token: t}),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
		grpc.WithUserAgent(fmt.Sprintf("Journaleux %s #%s", runtime.GOOS, git_version)),
	)
	return conn, err
}
