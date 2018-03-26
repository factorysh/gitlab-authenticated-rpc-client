package auth

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/skratchdot/open-golang/open"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/conf"
)

const (
	AUTH_CODE_URL = "gar.auth_code_url"
	SESSION_ID    = "gar.session_id"
	TOKEN         = "gar.oauth_token"
)

type Auth struct {
	Token     string
	SessionId string
	Conf      *conf.Conf
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	m := make(map[string]string)
	if a.Token != "" {
		m["authorization"] = a.Token
	}
	if a.SessionId != "" {
		m["session"] = a.SessionId
	}
	return m, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return true
}

func (a *Auth) AuthInterceptor(ctx context.Context, method string, req, resp interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var input string
	retry := 0
	for {
		md := metadata.Pairs()
		newOpts := append(opts, grpc.Trailer(&md))
		newCtx := context.WithValue(ctx, "gar.retry", retry)
		err := invoker(newCtx, method, req, resp, cc, newOpts...)
		// log.Printf("%#v", md)
		// if the server send a token, store it
		t, ok := md[TOKEN]
		if ok {
			a.SessionId = ""
			a.Token = t[0]
			err := a.Conf.SetToken(a.Token)
			if err != nil {
				log.Fatal(err)
			}
		}
		if err != nil {
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.Unauthenticated {
				s, ok := md[SESSION_ID]
				if ok {
					a.SessionId = s[0]
				}
				u, ok := md[AUTH_CODE_URL]
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
					retry++
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
