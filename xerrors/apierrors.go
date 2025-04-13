package xerrors

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrAPIStatusIllegal      = APIStatusIllegal("default status illegal")
	ErrAPIParamInvalid       = APIParamInvalid("default param invalid")
	ErrAPIPageParamInvalid   = APIPageParamInvalid("default page param invalid")
	ErrAPINotFound           = APINotFound("default not found")
	ErrAPIAlreadyExists      = APIAlreadyExists("default already exists")
	ErrAPIStateUpdateFailed  = APIStateUpdateFailed("default state update failed")
	ErrAPISessionIllegal     = APISessionIllegal("default session illegal")
	ErrAPISessionTimeout     = APISessionTimeout("default session timeout")
	ErrAPIAuthFailed         = APIAuthFailed("default auth failed")
	ErrAPIPlatformAuthFailed = APIPlatformAuthFailed("default platform auth failed")
	ErrAPICodecFailed        = APICodecFailed("default codec failed")
	ErrAPIDBFailed           = APIDBFailed("default db failed")
)

func APIStatusIllegal(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.Forbidden("STATUS_ILLEGAL", message)
}

func APIParamInvalid(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.BadRequest("PARAM_INVALID", message)
}

func APIPageParamInvalid(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.BadRequest("PAGE_PARAM_INVALID", message)
}

func APINotFound(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.NotFound("NOT_FOUND", message)
}

func APIAlreadyExists(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.Conflict("ALREADY_EXISTS", message)
}

func APIStateUpdateFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.Conflict("STATE_UPDATE_FAILED", message)
}

func APISessionIllegal(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.Unauthorized("SESSION_ILLEGAL", message)
}

func APISessionTimeout(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.Unauthorized("SESSION_TIMEOUT", message)
}

func APIAuthFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.Unauthorized("AUTH_FAILED", message)
}

func APIPlatformAuthFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.Unauthorized("PLATFORM_AUTH_FAILED", message)
}

func APICodecFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.InternalServer("CODEC_FAILED", message)
}

func APIDBFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}
	return errors.InternalServer("DB_FAILED", message)
}
