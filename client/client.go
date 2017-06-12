package client

import (
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/auth"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/version"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"runtime"
	"strings"
	"time"
)

const (
	port = ":50051"
)

func NewConn(domain string) (*grpc.ClientConn, error) {

	cfg := conf.NewConf("gar", domain)
	t, err := cfg.Token()
	if err != nil {
		return nil, errors.Wrap(err, "Can't get token")
	}

	// Set up a connection to the server.
	// doc https://godoc.org/google.golang.org/grpc#Dial
	conn, err := grpc.Dial(domain+port,
		grpc.WithPerRPCCredentials(&auth.IdAuth{Token: t}),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true, //FIXME don't do that on prod
		})),
		grpc.WithUnaryInterceptor(askForToken),
		grpc.WithUserAgent(fmt.Sprintf("GAR %s #%s", runtime.GOOS, version.GitVersion)),
		grpc.FailOnNonTempDialError(true),
		// set a timeout
		grpc.WithTimeout(4*time.Second),
		// block until sucess or failure (needed to set err correctly)
		grpc.WithBlock(),
	)
	return conn, err
}

func askForToken(ctx context.Context, method string, req, resp interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	md := metadata.Pairs()
	opts = append(opts, grpc.Trailer(&md))
	err := invoker(ctx, method, req, resp, cc, opts...)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.Unauthenticated {
			u, ok := md["gar.auth_code_url"]
			if ok {
				if !strings.HasPrefix(u[0], "https://") {
					log.Fatal("Bad url prefix, please ensure an https endpoint")
				}
				open.Run(u[0])
				return fmt.Errorf("Invalid session, please authenticate at %s\n", u[0])
			}
		}
		log.Fatalf("Can't hello: %v %v\n", err, md)
		return err
	}
	return nil
}
