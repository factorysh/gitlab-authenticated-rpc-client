package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/auth"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"runtime"
	"strings"
	"time"
)

/*
certPool can be nil or contains a private CA, for non public TLS chain
*/
func NewConn(domain string, certPool *x509.CertPool) (*grpc.ClientConn, error) {
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
	a := &auth.Auth{
		Token:     t,
		SessionId: "",
		Conf:      cfg,
	}
	options := []grpc.DialOption{
		grpc.WithPerRPCCredentials(a),
		grpc.WithUnaryInterceptor(a.AuthInterceptor),
		grpc.WithUserAgent(fmt.Sprintf("GAR %s #%s", runtime.GOOS, version.GitVersion)),
		grpc.FailOnNonTempDialError(true),
		// set a timeout
		grpc.WithTimeout(4 * time.Second),
		// block until sucess or failure (needed to set err correctly)
		grpc.WithBlock(),
	}
	if certPool != nil {
		dialer := func(address string, timeout time.Duration) (net.Conn, error) {
			return tls.Dial("tcp", address, &tls.Config{
				RootCAs: certPool,
			})
		}
		options = append(options, grpc.WithInsecure(), grpc.WithDialer(dialer))
	} else {
		options = append(options, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true, //FIXME don't do that on prod
		})))
	}
	conn, err := grpc.Dial(domain, options...)

	if err != nil {
		return conn, fmt.Errorf("Can't connect to %s, is the remote service up ?", domain)
	}

	return conn, err
}
