// Package xcontext provides context-related utilities for propagating
// and accessing metadata through the service chain, including user IDs,
// routing information, and status data.
package xcontext

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/pkg/errors"
	grpcmd "google.golang.org/grpc/metadata"
)

// Context is the context of the game
// Use the custom type for your constants
const (
	CtxSID         = "x-md-global-sid"     // Server ID is the ID for each server in multi-server games, or 0 for single-server games
	CtxUID         = "x-md-global-uid"     // User ID is the ID of the player.It is unique in the game.
	CtxOID         = "x-md-global-oid"     // Object ID for route the message to specific node which has the corresponding module and ID
	CtxColor       = "x-md-global-color"   // Color for route the message to specific node group
	CtxStatus      = "x-md-global-status"  // Status is the status of this connection
	CtxReferer     = "x-md-global-referer" // example: gate:10.0.1.31 or player:10.0.2.31
	CtxClientIP    = "x-md-global-client-ip"
	CtxGateReferer = "x-md-global-gate-referer" // example: 10.0.1.31:9100#10001
)

// Keys is a list of all context metadata keys used in the system.
var Keys = []string{CtxSID, CtxUID, CtxOID, CtxStatus, CtxColor, CtxReferer, CtxClientIP, CtxGateReferer}

func AppendToClientContext(ctx context.Context, kv ...string) context.Context {
	if len(kv)%2 != 0 {
		panic("append to client context: kv must be even")
	}

	mds := make(map[string][]string, len(kv)/2)

	for i := 0; i < len(kv); i += 2 {
		mds[kv[i]] = []string{kv[i+1]}
	}

	return metadata.MergeToClientContext(ctx, metadata.New(mds))
}

func AppendToServerContext(ctx context.Context, kv ...string) context.Context {
	if len(kv)%2 != 0 {
		panic("append to server context: kv must be even")
	}

	var md metadata.Metadata

	smd, ok := metadata.FromServerContext(ctx)
	if ok {
		md = metadata.New(smd)
	} else {
		md = metadata.New(make(map[string][]string, len(kv)/2))
	}

	for i := 0; i < len(kv); i += 2 {
		md[kv[i]] = []string{kv[i+1]}
	}

	return metadata.NewServerContext(ctx, md)
}

// Color retrieves the color information from the server context.
func Color(ctx context.Context) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(string(CtxColor))
	}

	return ""
}

// UID retrieves the user ID from the server context.
// Returns an error if the context doesn't contain metadata or if the ID is not a valid int64.
func UID(ctx context.Context) (int64, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0, errors.New("metadata not in context")
	}

	v := md.Get(string(CtxUID))

	id, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "uid must be int64, uid=%s", v)
	}

	return id, nil
}

func UIDOrZero(ctx context.Context) int64 {
	uid, err := UID(ctx)
	if err != nil {
		return 0
	}

	return uid
}

// OID retrieves the object ID from the server context.
// Returns an error if the context doesn't contain metadata or if the ID is not a valid int64.
func OID(ctx context.Context) (int64, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0, errors.New("metadata not in context")
	}

	v := md.Get(string(CtxOID))

	id, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "oid must be int64, oid=%s", v)
	}

	return id, nil
}

func OIDOrZero(ctx context.Context) int64 {
	oid, err := OID(ctx)
	if err != nil {
		return 0
	}

	return oid
}

// SID retrieves the server ID from the server context.
// Returns an error if the context doesn't contain metadata or if the ID is not a valid int64.
func SID(ctx context.Context) (int64, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0, errors.New("metadata not in context")
	}

	v := md.Get(string(CtxSID))

	id, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "sid must be int64, sid=%s", v)
	}

	return id, nil
}

func SIDOrZero(ctx context.Context) int64 {
	sid, err := SID(ctx)
	if err != nil {
		return 0
	}

	return sid
}

// Status retrieves the status information from the server context.
// Returns 0 if the context doesn't contain metadata or if the status is not a valid int64.
func Status(ctx context.Context) int64 {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0
	}

	v := md.Get(string(CtxStatus))

	status, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}

	return status
}

// ClientIP retrieves the client IP address from the server context.
func ClientIP(ctx context.Context) string {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return ""
	}

	return md.Get(string(CtxClientIP))
}

// GateReferer retrieves the gate server reference information from the server context.
func GateReferer(ctx context.Context) string {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return ""
	}

	return md.Get(string(CtxGateReferer))
}

func ColorFromOutgoingContext(ctx context.Context) string {
	if md, ok := grpcmd.FromOutgoingContext(ctx); ok {
		if v := md.Get(string(CtxColor)); len(v) > 0 {
			return v[0]
		}
	}

	return ""
}

func OIDFromOutgoingContext(ctx context.Context) (int64, error) {
	md, ok := grpcmd.FromOutgoingContext(ctx)
	if !ok {
		return 0, errors.New("metadata not in context")
	}

	if v := md.Get(string(CtxOID)); len(v) > 0 {
		id, err := strconv.ParseInt(v[0], 10, 64)
		if err != nil {
			return 0, errors.Wrapf(err, "oid must be int64, oid=%s", v[0])
		}

		return id, nil
	}

	return 0, errors.New("oid not found")
}
