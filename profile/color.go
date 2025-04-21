package profile

import "strings"

// Color is the color of the service, the message will be routed to the corresponding color node
const (
	ColorLocal = "local"
)

// IsLocal checks if the current service color is local.
// Returns true if the service is running in local mode.
func IsLocal() bool {
	return strings.ToLower(_color) == ColorLocal
}
