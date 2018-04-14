package auth

import (
	"fmt"
	"runtime"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/prometheus/common/log"
	"github.com/skratchdot/open-golang/open"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/version"
	_auth "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc_auth"
)

// Auth client
type Auth struct {
	SessionID string
	Conf      *conf.Conf
	client    *grpc.ClientConn
}

func New(cfg *conf.Conf) *Auth {
	return &Auth{
		Conf: cfg,
	}
}

func (a *Auth) cliencConn() (*grpc.ClientConn, error) {
	if a.client != nil {
		return a.client, nil
	}
	options := []grpc.DialOption{
		grpc.WithUserAgent(fmt.Sprintf("GAR %s #%s", runtime.GOOS, version.GitVersion)),
		grpc.FailOnNonTempDialError(true),
		// set a timeout
		grpc.WithTimeout(4 * time.Second),
		// block until sucess or failure (needed to set err correctly)
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
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
	if jwt == "" {
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
	if st.Code() != codes.Unauthenticated {
		return rpcErr
	}
	// Handle unauthenticated error
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
