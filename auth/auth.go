package auth

import (
	"crypto/x509"

	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/dial"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/prometheus/common/log"
	"github.com/skratchdot/open-golang/open"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	_auth "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc_auth"
)

// Auth client
type Auth struct {
	SessionID           string
	Conf                *conf.Conf
	CertPool            *x509.CertPool
	client              *grpc.ClientConn
	TryToAuthWithGitlab bool
}

func New(cfg *conf.Conf, cert *x509.CertPool) *Auth {
	return &Auth{
		Conf:                cfg,
		CertPool:            cert,
		TryToAuthWithGitlab: true,
	}
}

func (a *Auth) cliencConn() (*grpc.ClientConn, error) {
	if a.client != nil {
		return a.client, nil
	}
	options := dial.ClientDialOptions(a.CertPool)

	// TODO domain can come from an header
	conn, err := grpc.Dial(a.Conf.Domain, options...)
	if err != nil {
		return nil, err
	}
	a.client = conn
	return conn, nil
}

// GetRequestMetadata gets the current request metadata
// Implements https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials
func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	log.Info("GetRequestMetadata uri:", uri)
	t, err := a.Conf.GetToken()
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"authorization": "bearer " + t,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security.
// Implements https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials
func (a *Auth) RequireTransportSecurity() bool {
	return true
}

// AuthInterceptor intercepts the execution of a unary RPC on the client
// Implements https://godoc.org/google.golang.org/grpc#UnaryClientInterceptor
func (a *Auth) AuthInterceptor(ctx context.Context, method string, req, resp interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	jwt, err := a.Conf.GetToken()
	if err != nil {
		return err
	}
	var ctx2 context.Context
	if jwt == "" && a.TryToAuthWithGitlab {
		ctx2, err = a.authDance(ctx)
		if err != nil {
			return err
		}
	} else {
		ctx2 = ctx
	}
	rpcErr := invoker(ctx2, method, req, resp, cc, opts...)
	if rpcErr == nil {
		return nil
	}
	st, ok := status.FromError(rpcErr)
	if !ok { // It's not an http error
		return rpcErr
	}
	if (st.Code() != codes.Unauthenticated) || !a.TryToAuthWithGitlab {
		return rpcErr
	}
	// Handle unauthenticated error and try to authenticate with Gitlab
	ctx2, err = a.authDance(ctx)
	if err != nil {
		return err
	}

	// FIXME set token in the header
	return invoker(ctx2, method, req, resp, cc, opts...)
}

func (a *Auth) authorize(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx,
		"authorization", "bearer "+token)
}

func (a *Auth) authDance(ctx context.Context) (context.Context, error) {
	log.Info("Auth dance")
	cc, err := a.cliencConn()
	if err != nil {
		return ctx, err
	}
	aa := _auth.NewAuthClient(cc)
	// FIXME a.Client == nil
	authCtx := context.Background()
	authinfo, err := aa.Bootstrap(authCtx, &empty.Empty{})
	if err != nil {
		return ctx, err
	}
	open.Run(authinfo.Url)
	j, err := aa.Authenticate(authCtx, &_auth.Token{
		Token: authinfo.Token,
	})
	if err != nil {
		return ctx, err
	}
	err = a.Conf.SetToken(j.JWT)
	return a.authorize(ctx, authinfo.Token), err
}
