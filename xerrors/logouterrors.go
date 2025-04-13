package xerrors

import "github.com/pkg/errors"

// User logout errors
var (
	ErrLogoutFromUser         = errors.New("user logout")
	ErrLogoutBanned           = errors.New("banned")
	ErrLogoutKickOut          = errors.New("kick out")
	ErrLogoutConflictingLogin = errors.New("conflicting login")
	ErrLogoutMainTunnelClosed = errors.New("main tunnel closed")
)

func IsLogoutError(err error) bool {
	return errors.Is(err, ErrLogoutFromUser) ||
		errors.Is(err, ErrLogoutBanned) ||
		errors.Is(err, ErrLogoutKickOut) ||
		errors.Is(err, ErrLogoutConflictingLogin) ||
		errors.Is(err, ErrLogoutMainTunnelClosed)
}
