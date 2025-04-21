package xerrors

import "github.com/go-pantheon/fabrica-util/errors"

// User logout errors
var (
	// ErrLogoutFromUser is returned when a user actively logs out.
	ErrLogoutFromUser = errors.New("user logout")
	// ErrLogoutBanned is returned when a user is banned from the service.
	ErrLogoutBanned = errors.New("banned")
	// ErrLogoutKickOut is returned when a user is forcibly disconnected.
	ErrLogoutKickOut = errors.New("kick out")
	// ErrLogoutConflictingLogin is returned when the same user logs in from another location.
	ErrLogoutConflictingLogin = errors.New("conflicting login")
	// ErrLogoutMainTunnelClosed is returned when the main communication tunnel is closed.
	ErrLogoutMainTunnelClosed = errors.New("main tunnel closed")
)

// IsLogoutError checks if the provided error is any type of logout error.
// Returns true if the error is one of the predefined logout errors.
func IsLogoutError(err error) bool {
	return errors.Is(err, ErrLogoutFromUser) ||
		errors.Is(err, ErrLogoutBanned) ||
		errors.Is(err, ErrLogoutKickOut) ||
		errors.Is(err, ErrLogoutConflictingLogin) ||
		errors.Is(err, ErrLogoutMainTunnelClosed)
}
