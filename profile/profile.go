package profile

import "strings"

// Profile is the running model of the service
const (
	ProfileDev  = "dev"
	ProfileTest = "test"
	ProfileProd = "prod"
)

// IsDev checks if the current profile is development.
// Returns true if the service is running in development mode.
func IsDev() bool {
	return IsDevStr(_profile)
}

// IsDevStr checks if the given profile string is the development profile.
// Returns true if the profile string matches the development profile.
func IsDevStr(profile string) bool {
	return strings.ToLower(profile) == ProfileDev
}

// IsTestStr checks if the given profile string is the test profile.
// Returns true if the profile string matches the test profile.
func IsTestStr(profile string) bool {
	return strings.ToLower(profile) == ProfileTest
}

// IsProdStr checks if the given profile string is the production profile.
// Returns true if the profile string matches the production profile.
func IsProdStr(profile string) bool {
	return strings.ToLower(profile) == ProfileProd
}
