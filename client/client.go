package client

import (
	"crypto/x509"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/auth"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/version"
	_rpc "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	Domain         string
	Client         *grpc.ClientConn
	Token          string
	CertPool       *x509.CertPool
	AuthWithGitlab bool
}

func New(domain string) *Client {
	if len(strings.Split(domain, ":")) == 1 {
		domain = domain + ":50051"
	}
	return &Client{
		Domain:         domain,
		AuthWithGitlab: true,
	}
}

func (c *Client) ClientConn() (*grpc.ClientConn, error) {

	cfg := conf.NewConf("gar", c.Domain)
	var t string
	if c.Token != "" {
		t = c.Token
	} else {
		var err error
		t, err = cfg.GetToken()
		if err != nil {
			return nil, errors.Wrap(err, "Can't get token")
		}
	}

	log.WithFields(log.Fields{
		"token":          t,
		"domain":         c.Domain,
		"with_your_pool": c.CertPool != nil,
	}).Info("Connecting")

	// Set up a connection to the server.
	// doc https://godoc.org/google.golang.org/grpc#Dial
	a := &auth.Auth{
		Token:     t,
		SessionID: "",
		Conf:      cfg,
	}
	options := []grpc.DialOption{
		grpc.WithPerRPCCredentials(a),
		grpc.WithUserAgent(fmt.Sprintf("GAR %s #%s", runtime.GOOS, version.GitVersion)),
		grpc.FailOnNonTempDialError(true),
		// set a timeout
		grpc.WithTimeout(4 * time.Second),
		// block until sucess or failure (needed to set err correctly)
		grpc.WithBlock(),
		grpc.WithTransportCredentials(
			credentials.NewClientTLSFromCert(c.CertPool, "")),
	}
	if c.AuthWithGitlab {
		options = append(options,
			grpc.WithUnaryInterceptor(a.AuthInterceptor),
		)
	}
	conn, err := grpc.Dial(c.Domain, options...)

	if err != nil {
		// FIXME better error handling : try TCP connect, TLS, and after grpc stuff
		return nil, fmt.Errorf("Can't connect to %s, is the remote service up ? %s", c.Domain, err)
	}
	return conn, err
}

func (c *Client) NewGitlabClient() (_rpc.GitlabClient, error) {
	cc, err := c.ClientConn()
	if err != nil {
		return nil, err
	}
	return _rpc.NewGitlabClient(cc), nil
}
