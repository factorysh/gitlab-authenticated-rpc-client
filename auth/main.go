package auth

import (
	"golang.org/x/net/context"
)

type IdAuth struct {
	Token string
}

func (i *IdAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Id " + i.Token,
	}, nil
}

func (i *IdAuth) RequireTransportSecurity() bool {
	return true
}
