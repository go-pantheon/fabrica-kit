package profile

import (
	"net/url"
	"strings"
)

var (
	_profile      string
	_color        string
	_version      string
	_grpcEndpoint string
	_nodeName     string
	_zone         uint32
)

// Init initializes the profile settings with the given parameters.
// It sets up the environment profile, color, zone, version, node name, and gRPC endpoint.
func Init(profile, color string, zone uint32, version string, nodeName string, gRPCEndpoint *url.URL) {
	_profile = profile
	_color = color
	_version = version
	_grpcEndpoint = strings.Replace(gRPCEndpoint.String(), "grpc://", "", 1)
	_nodeName = nodeName
	_zone = zone
}

// Profile returns the current environment profile.
func Profile() string {
	return _profile
}

// Color returns the current deployment color.
func Color() string {
	return _color
}

// Version returns the current version string.
func Version() string {
	return _version
}

// NodeName returns the name of the current node.
func NodeName() string {
	return _nodeName
}

// GRPCEndpoint returns the gRPC endpoint string.
func GRPCEndpoint() string {
	return _grpcEndpoint
}

// Zone returns the current zone identifier.
func Zone() uint32 {
	return _zone
}
