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

var (
	ErrAdminPermissionReason   = "permission error"
	ErrAdminQueryFailedReason  = "query execute failed"
	ErrAdminUpdateFailedReason = "update failed"
	ErrAdminParamReason        = "param error"
	ErrAdminConditionReason    = "condition error"
)
