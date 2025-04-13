package xerrors

import (
	"errors"
)

// Route table errors
var (
	ErrRouteTableNotFound = errors.New("route table not found")
)

// Tunnel errors
var (
	ErrTunnelStopped = errors.New("tunnel stopped")
)

// DB errors
var (
	ErrDBRecordNotFound    = errors.New("record not found")
	ErrDBRecordExists      = errors.New("record exists")
	ErrDBRecordVersion     = errors.New("record version error")
	ErrDBRecordNotAffected = errors.New("record update error")
	ErrDBRecordType        = errors.New("record type error")
	ErrDBProtoEncode       = errors.New("db proto encode error")
	ErrDBProtoDecode       = errors.New("db proto decode error")
)
