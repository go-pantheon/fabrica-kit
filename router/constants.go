// Package router provides routing capabilities for service discovery
// and network communication in distributed systems.
package router

import "time"

// AppTunnelChangeTimeout defines the timeout duration for tunnel connection changes.
const (
	AppTunnelChangeTimeout = time.Second * 3
	HolderCacheTimeout     = time.Second * 5
	AsyncRouteTableTimeout = time.Second * 1
)
