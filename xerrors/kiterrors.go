// Package xerrors provides standardized error types and error handling utilities
// for application-specific errors, including API, database, and routing errors.
package xerrors

import (
	"context"
	"io"

	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/fabrica-util/xsync"
)

// Route table errors
var (
	// ErrRouteTableNotFound is returned when a requested route table entry cannot be found.
	ErrRouteTableNotFound     = errors.New("route table not found")
	ErrRouteTableValueNotSame = errors.New("route table not same")
)

var (
	ErrHandlerNotFound = errors.New("handler not found")
)

// Tunnel errors
var (
	// ErrTunnelStopped is returned when operations are attempted on a stopped tunnel.
	ErrTunnelStopped = errors.New("tunnel stopped")
	ErrLifeStopped   = errors.New("life stopped")
)

// DB errors
var (
	// ErrDBRecordNotFound is returned when a requested database record does not exist.
	ErrDBRecordNotFound = errors.New("record not found")
	// ErrDBRecordExists is returned when attempting to create a record that already exists.
	ErrDBRecordExists = errors.New("record exists")
	// ErrDBRecordVersion is returned when there's a version conflict in optimistic locking.
	ErrDBRecordVersion = errors.New("record version error")
	// ErrDBRecordNotAffected is returned when a database update operation didn't affect any records.
	ErrDBRecordNotAffected = errors.New("record update error")
	// ErrDBRecordType is returned when there's a type mismatch in database operations.
	ErrDBRecordType = errors.New("record type error")
	// ErrDBProtoEncode is returned when protobuf encoding fails for a database record.
	ErrDBProtoEncode = errors.New("db proto encode error")
	// ErrDBProtoDecode is returned when protobuf decoding fails for a database record.
	ErrDBProtoDecode = errors.New("db proto decode error")
)

func IsUnlogErr(err error) bool {
	return errors.Is(err, xsync.ErrStopByTrigger) || IsEOFError(err) || IsCancelError(err) || IsLogoutError(err)
}

func IsEOFError(err error) bool {
	return errors.Is(err, io.EOF)
}

func IsCancelError(err error) bool {
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}
