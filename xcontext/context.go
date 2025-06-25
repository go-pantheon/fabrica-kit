// Package xcontext provides context-related utilities for propagating
// and accessing metadata through the service chain, including user IDs,
// routing information, and status data.
package xcontext

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/pkg/errors"
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

// SetColor adds the color information to the client context.
func SetColor(ctx context.Context, color string) context.Context {
	return metadata.AppendToClientContext(ctx, string(CtxColor), color)
}

// Color retrieves the color information from the server context.
func Color(ctx context.Context) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(CtxColor)
	}

	return ""
}

// SetUID adds the user ID to the client context.
func SetUID(ctx context.Context, id int64) context.Context {
	return metadata.AppendToClientContext(ctx, CtxUID, strconv.FormatInt(id, 10))
}

// UID retrieves the user ID from the server context.
// Returns an error if the context doesn't contain metadata or if the ID is not a valid int64.
func UID(ctx context.Context) (int64, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0, errors.New("metadata not in context")
	}

	str := md.Get(CtxUID)
	id, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return 0, errors.Wrapf(err, "uid must be int64, uid=%s", str)
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

// SetOID adds the object ID to the client context.
func SetOID(ctx context.Context, id int64) context.Context {
	return metadata.AppendToClientContext(ctx, CtxOID, strconv.FormatInt(id, 10))
}

// OID retrieves the object ID from the server context.
// Returns an error if the context doesn't contain metadata or if the ID is not a valid int64.
func OID(ctx context.Context) (int64, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0, errors.New("metadata not in context")
	}

	str := md.Get(CtxOID)
	id, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return 0, errors.Wrapf(err, "oid must be int64, oid=%s", str)
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

// SetSID adds the server ID to the client context.
func SetSID(ctx context.Context, id int64) context.Context {
	return metadata.AppendToClientContext(ctx, CtxSID, strconv.FormatInt(id, 10))
}

// SID retrieves the server ID from the server context.
// Returns an error if the context doesn't contain metadata or if the ID is not a valid int64.
func SID(ctx context.Context) (int64, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0, errors.New("metadata not in context")
	}

	str := md.Get(CtxSID)
	id, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return 0, errors.Wrapf(err, "sid must be int64, sid=%s", str)
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

// SetStatus adds the status information to the client context.
// If status is 0, the original context is returned without modification.
func SetStatus(ctx context.Context, status int64) context.Context {
	if status == 0 {
		return ctx
	}

	return metadata.AppendToClientContext(ctx, CtxStatus, strconv.FormatInt(status, 10))
}

// Status retrieves the status information from the server context.
// Returns 0 if the context doesn't contain metadata or if the status is not a valid int64.
func Status(ctx context.Context) int64 {
	if md, ok := metadata.FromServerContext(ctx); ok {
		v := md.Get(CtxStatus)
		status, err := strconv.ParseInt(v, 10, 64)

		if err != nil {
			log.Errorf("status must be int64, status=%s", v)
			return 0
		}

		return status
	}

	return 0
}

// SetClientIP adds the client IP address to the client context.
// If the IP is empty, the original context is returned without modification.
func SetClientIP(ctx context.Context, ip string) context.Context {
	if len(ip) == 0 {
		return ctx
	}

	return metadata.AppendToClientContext(ctx, CtxClientIP, strings.Split(ip, ":")[0])
}

// ClientIP retrieves the client IP address from the server context.
func ClientIP(ctx context.Context) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(CtxClientIP)
	}

	return ""
}

// SetGateReferer adds the gate server reference information to the client context.
// If the server string is empty, the original context is returned without modification.
func SetGateReferer(ctx context.Context, server string, wid uint64) context.Context {
	if len(server) == 0 {
		return ctx
	}

	return metadata.AppendToClientContext(ctx, CtxGateReferer, fmt.Sprintf("%s#%d", server, wid))
}

// GateReferer retrieves the gate server reference information from the server context.
func GateReferer(ctx context.Context) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(CtxGateReferer)
	}

	return ""
}
