package conf

import (
	"crypto/x509"
	"errors"
	"io/ioutil"
)

/*
path to CA file in pem format
*/
func GetCA(path string) (*x509.CertPool, error) {
	if path == "" {
		return nil, nil
	}
	pool := x509.NewCertPool()
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ok := pool.AppendCertsFromPEM(raw)
	if !ok {
		return nil, errors.New("failed to parse root certificate")
	}
	return pool, nil
}
