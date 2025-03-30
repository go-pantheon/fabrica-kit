package xerrors

import (
	"errors"

	kerrors "github.com/go-kratos/kratos/v2/errors"
)

// Route table errors
var (
	ErrRouteTableNotFound = errors.New("route table not found")
)

// Tunnel errors
var (
	ErrTunnelStopped = errors.New("tunnel stopped")
)

// User logout errors
var (
	ErrLogoutFromUser         = errors.New("user logout")
	ErrLogoutBanned           = errors.New("banned")
	ErrLogoutKickOut          = errors.New("kick out")
	ErrLogoutConflictingLogin = errors.New("conflicting login")
	ErrLogoutMainTunnelClosed = errors.New("main tunnel closed")
)

// Client request errors
var (
	ErrAPIServerErr       = kerrors.InternalServer("server error", "please try again later")
	ErrAPIStatusIllegal   = kerrors.Forbidden("request forbidden", "request status error")
	ErrAPISessionErr      = kerrors.Unauthorized("unauthorized", "session error")
	ErrAPIPasswordInvalid = kerrors.Unauthorized("unauthorized", "password error")
	ErrAPIRequestInvalid  = kerrors.BadRequest("illegal request", "parameter error")
	ErrAPIPlatformInvalid = kerrors.BadRequest("illegal request", "platform id error")
)

// DB errors
var (
	ErrDBProtoEncode       = errors.New("db proto encode error")
	ErrDBProtoDecode       = errors.New("db proto decode error")
	ErrDBRecordNotFound    = errors.New("record not found")
	ErrDBRecordExists      = errors.New("record exists")
	ErrDBRecordVersion     = errors.New("record version error")
	ErrDBRecordNotAffected = errors.New("record update error")
	ErrDBRecordType        = errors.New("record type error")
)
