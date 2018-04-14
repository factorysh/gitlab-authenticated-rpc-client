package client

import (
	"crypto/x509"
	"fmt"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/auth"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/version"
	_rpc "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
	_auth "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc_auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Client is the high level client
type Client struct {
	Domain         string
	client         *grpc.ClientConn
	CertPool       *x509.CertPool
	AuthWithGitlab bool
	Conf           *conf.Conf
}

// New client
func New(domain string) *Client {
	if len(strings.Split(domain, ":")) == 1 {
		domain = domain + ":50051"
	}
	return &Client{
		Domain:         domain,
		AuthWithGitlab: true,
		Conf:           conf.NewConf("gar", domain),
	}
}

// ClientConn is the grpc client connection
func (c *Client) ClientConn(wirthRPCCredential bool) (*grpc.ClientConn, error) {
	log.WithFields(log.Fields{
		"domain":         c.Domain,
		"with_your_pool": c.CertPool != nil,
	}).Info("Connecting")

	// Set up a connection to the server.
	// doc https://godoc.org/google.golang.org/grpc#Dial
	options := []grpc.DialOption{
		grpc.WithUserAgent(fmt.Sprintf("GAR %s #%s", runtime.GOOS, version.GitVersion)),
		grpc.FailOnNonTempDialError(true),
		// set a timeout
		grpc.WithTimeout(15 * time.Second),
		// block until sucess or failure (needed to set err correctly)
		grpc.WithBlock(),
		grpc.WithTransportCredentials(
			credentials.NewClientTLSFromCert(c.CertPool, "")),
	}
	if wirthRPCCredential {
		a := auth.New(c.Conf)
		options = append(options, grpc.WithPerRPCCredentials(a))
		if c.AuthWithGitlab {
			options = append(options,
				grpc.WithUnaryInterceptor(a.AuthInterceptor),
			)
		}
	}
	conn, err := grpc.Dial(c.Domain, options...)

	if err != nil {
		// FIXME better error handling : try TCP connect, TLS, and after grpc stuff
		return nil, fmt.Errorf("Can't connect to %s, is the remote service up ? %s", c.Domain, err)
	}
	return conn, err
}

// NewGitlabClient : new grpc Gitlab client instance
func (c *Client) NewGitlabClient() (_rpc.GitlabClient, error) {
	cc, err := c.ClientConn(true)
	if err != nil {
		return nil, err
	}
	return _rpc.NewGitlabClient(cc), nil
}

// NewAuthClient : new grpc Auth client instance
func (c *Client) NewAuthClient() (_auth.AuthClient, error) {
	cc, err := c.ClientConn(false)
	if err != nil {
		return nil, err
	}
	return _auth.NewAuthClient(cc), nil
}
