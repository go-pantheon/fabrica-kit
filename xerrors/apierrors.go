package xerrors

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrAPIStatusIllegal is a predefined error for illegal status conditions.
	ErrAPIStatusIllegal = APIStatusIllegal("default status illegal")
	// ErrAPIParamInvalid is a predefined error for invalid parameters.
	ErrAPIParamInvalid = APIParamInvalid("default param invalid")
	// ErrAPIPageParamInvalid is a predefined error for invalid pagination parameters.
	ErrAPIPageParamInvalid = APIPageParamInvalid("default page param invalid")
	// ErrAPINotFound is a predefined error for resources that cannot be found.
	ErrAPINotFound = APINotFound("default not found")
	// ErrAPIAlreadyExists is a predefined error for resources that already exist.
	ErrAPIAlreadyExists = APIAlreadyExists("default already exists")
	// ErrAPIStateUpdateFailed is a predefined error for state update failures.
	ErrAPIStateUpdateFailed = APIStateUpdateFailed("default state update failed")
	// ErrAPISessionIllegal is a predefined error for illegal session conditions.
	ErrAPISessionIllegal = APISessionIllegal("default session illegal")
	// ErrAPISessionTimeout is a predefined error for session timeouts.
	ErrAPISessionTimeout = APISessionTimeout("default session timeout")
	// ErrAPIAuthFailed is a predefined error for authentication failures.
	ErrAPIAuthFailed = APIAuthFailed("default auth failed")
	// ErrAPIPlatformAuthFailed is a predefined error for platform authentication failures.
	ErrAPIPlatformAuthFailed = APIPlatformAuthFailed("default platform auth failed")
	// ErrAPICodecFailed is a predefined error for codec failures.
	ErrAPICodecFailed = APICodecFailed("default codec failed")
	// ErrAPIDBFailed is a predefined error for database operation failures.
	ErrAPIDBFailed = APIDBFailed("default db failed")
	// ErrAPIDBNoAffected is a predefined error for database operations that didn't affect any records.
	ErrAPIDBNoAffected = APIDBFailed("default db no affected")
)

// APIStatusIllegal creates a forbidden error with a "STATUS_ILLEGAL" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APIStatusIllegal(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.Forbidden("STATUS_ILLEGAL", message)
}

// APIParamInvalid creates a bad request error with a "PARAM_INVALID" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APIParamInvalid(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.BadRequest("PARAM_INVALID", message)
}

// APIPageParamInvalid creates a bad request error with a "PAGE_PARAM_INVALID" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APIPageParamInvalid(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.BadRequest("PAGE_PARAM_INVALID", message)
}

// APINotFound creates a not found error with a "NOT_FOUND" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APINotFound(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.NotFound("NOT_FOUND", message)
}

// APIAlreadyExists creates a conflict error with an "ALREADY_EXISTS" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APIAlreadyExists(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.Conflict("ALREADY_EXISTS", message)
}

// APIStateUpdateFailed creates a conflict error with a "STATE_UPDATE_FAILED" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APIStateUpdateFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.Conflict("STATE_UPDATE_FAILED", message)
}

// APISessionIllegal creates an unauthorized error with a "SESSION_ILLEGAL" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APISessionIllegal(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.Unauthorized("SESSION_ILLEGAL", message)
}

// APISessionTimeout creates an unauthorized error with a "SESSION_TIMEOUT" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APISessionTimeout(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.Unauthorized("SESSION_TIMEOUT", message)
}

// APIAuthFailed creates an unauthorized error with an "AUTH_FAILED" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APIAuthFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.Unauthorized("AUTH_FAILED", message)
}

// APIPlatformAuthFailed creates an unauthorized error with a "PLATFORM_AUTH_FAILED" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APIPlatformAuthFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.Unauthorized("PLATFORM_AUTH_FAILED", message)
}

// APICodecFailed creates an internal server error with a "CODEC_FAILED" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APICodecFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.InternalServer("CODEC_FAILED", message)
}

// APIDBFailed creates an internal server error with a "DB_FAILED" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APIDBFailed(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.InternalServer("DB_FAILED", message)
}

// APIDBNoAffected creates a conflict error with a "DB_NO_AFFECTED" reason code.
// The message can include format specifiers that will be replaced by the provided arguments.
func APIDBNoAffected(message string, a ...any) *errors.Error {
	if len(a) > 0 {
		message = fmt.Sprintf(message, a...)
	}

	return errors.Conflict("DB_NO_AFFECTED", message)
}
