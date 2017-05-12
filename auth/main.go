package auth

import (
	"golang.org/x/net/context"
)

type JWTAuth struct {
	Token string
}

func (j *JWTAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + j.Token,
	}, nil
}

func (j *JWTAuth) RequireTransportSecurity() bool {
	return true
}
