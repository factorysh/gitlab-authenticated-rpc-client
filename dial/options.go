package dial

import (
	"crypto/x509"
	"fmt"
	"runtime"
	"time"

	"github.com/factorysh/gitlab-authenticated-rpc-client/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Timeout is client timeout
const Timeout time.Duration = 15 * time.Second

// ClientDialOptions return base client options
func ClientDialOptions(cert *x509.CertPool) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithUserAgent(fmt.Sprintf("GAR %s #%s", runtime.GOOS, version.GitVersion)),
		grpc.FailOnNonTempDialError(true),
		// set a timeout
		grpc.WithTimeout(Timeout),
		// block until success or failure (needed to set err correctly)
		grpc.WithBlock(),
		grpc.WithTransportCredentials(
			credentials.NewClientTLSFromCert(cert, "")),
	}
}
