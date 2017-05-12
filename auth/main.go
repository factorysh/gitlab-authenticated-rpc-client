package auth

import (
	"encoding/base64"
	"encoding/json"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type JWTAuth struct {
	Token *oauth2.Token
}

func (j *JWTAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	data, err := json.Marshal(j.Token)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"authorization": "Bearer " + base64.StdEncoding.EncodeToString(data),
	}, nil
}

func (j *JWTAuth) RequireTransportSecurity() bool {
	return true
}
