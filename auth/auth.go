package auth

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
	garMetadata "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/metadata"
)

type Auth struct {
	Token     string
	SessionID string
	Conf      *conf.Conf
}

// GetRequestMetadata gets the current request metadata
// Implements https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials
func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	m := make(map[string]string)
	if a.Token != "" {
		m["authorization"] = a.Token
	}
	if a.SessionID != "" {
		m["session"] = a.SessionID
	}
	return m, nil
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
	var input string
	retry := 0
	for {
		md := metadata.Pairs()
		log.WithFields(log.Fields{
			"metadata": md,
		}).Info("AuthInterceptor")
		newOpts := append(opts, grpc.Trailer(&md))
		newCtx := context.WithValue(ctx, "gar.retry", retry)
		err := invoker(newCtx, method, req, resp, cc, newOpts...)
		//FIXME handle this error
		// log.Printf("%#v", md)
		// if the server send a token, store it
		t, ok := md[garMetadata.Token]
		if ok {
			a.SessionID = ""
			a.Token = t[0]
			err := a.Conf.SetToken(a.Token)
			if err != nil {
				log.Fatal(err)
			}
		}
		if err != nil {
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.Unauthenticated {
				s, ok := md[garMetadata.SessionID]
				if ok {
					a.SessionID = s[0]
				}
				u, ok := md[garMetadata.AuthCodeURL]
				if ok {
					if !strings.HasPrefix(u[0], "https://") {
						log.Fatal("Bad url prefix, please ensure an https endpoint")
					}
					open.Run(u[0])
					fmt.Printf("Invalid session:")
					fmt.Printf("\n\tIn order to generate a new session, please authenticate at:")
					fmt.Printf("\n\n\t%s", u[0])
					fmt.Printf("\n\n\tThen press Enter\n")
					fmt.Scanln(&input)
					//FIXME where this input is handled?
					retry++
				} else {
					log.Fatal("Server didn't send authentication url")
				}
			} else {
				log.Fatalf("Can't hello: %v %v\n", err, md)
				return err
			}
		} else {
			return nil
		}
	}
	return nil
}
