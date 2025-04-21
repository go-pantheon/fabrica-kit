package profile

import "time"

const (
	// ClientTimeout defines the default timeout for HTTP client requests.
	ClientTimeout = 10 * time.Second
	// ClientMaxIdleConns defines the maximum number of idle connections.
	ClientMaxIdleConns = 100
	// ClientMaxIdleConnsPerHost defines the maximum idle connections per host.
	ClientMaxIdleConnsPerHost = 100
	// ClientIdleConnTimeout defines how long an idle connection is kept in the pool.
	ClientIdleConnTimeout = 90 * time.Second
	// ClientTLSHandshakeTimeout defines the timeout for TLS handshake.
	ClientTLSHandshakeTimeout = 5 * time.Second
)
