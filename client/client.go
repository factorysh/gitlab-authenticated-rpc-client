package client

import (
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/auth"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"runtime"
	"strings"
	"time"
)

func NewConn(domain string) (*grpc.ClientConn, error) {
	if len(strings.Split(domain, ":")) == 1 {
		domain = domain + ":50051"
	}

	cfg := conf.NewConf("gar", domain)
	t, err := cfg.GetToken()
	if err != nil {
		return nil, errors.Wrap(err, "Can't get token")
	}

	// Set up a connection to the server.
	// doc https://godoc.org/google.golang.org/grpc#Dial
	a := &auth.Auth{Token: t, SessionId: "", Conf: cfg}
	conn, err := grpc.Dial(domain,
		grpc.WithPerRPCCredentials(a),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true, //FIXME don't do that on prod
		})),
		grpc.WithUnaryInterceptor(a.AuthInterceptor),
		grpc.WithUserAgent(fmt.Sprintf("GAR %s #%s", runtime.GOOS, version.GitVersion)),
		grpc.FailOnNonTempDialError(true),
		// set a timeout
		grpc.WithTimeout(4*time.Second),
		// block until sucess or failure (needed to set err correctly)
		grpc.WithBlock(),
	)

	if err != nil {
		return conn, fmt.Errorf("Can't connect to %s, is the remote service up ?", domain)
	}

	return conn, err
}
