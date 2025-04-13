package profile

import "time"

const (
	ClientTimeout             = 10 * time.Second
	ClientMaxIdleConns        = 100
	ClientMaxIdleConnsPerHost = 100
	ClientIdleConnTimeout     = 90 * time.Second
	ClientTLSHandshakeTimeout = 5 * time.Second
)
