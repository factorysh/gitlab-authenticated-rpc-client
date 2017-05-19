package auth

import (
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

type IdAuth struct {
	Token *uuid.UUID
}

func (i *IdAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Id " + i.Token.String(),
	}, nil
}

func (i *IdAuth) RequireTransportSecurity() bool {
	return true
}
