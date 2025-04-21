// Package version provides utilities for parsing and comparing version strings
// to ensure compatibility and properly handle versioned resources.
package version

import (
	"strconv"
	"strings"
)

// GetSubVersion parses a version string in the format "zone-vX.Y" and extracts its components.
// Returns:
//   - az: the availability zone or region part (before the hyphen)
//   - sv: a slice of integers representing the major and minor version numbers
//   - isRelease: whether the version string is properly formatted as a release version
//
// Returns empty values and false for isRelease if the version string is invalid.
func GetSubVersion(v string) (az string, sv []int64, isRelease bool) {
	if len(v) == 0 {
		return "", nil, false
	}

	ss := strings.Split(v, "-")
	if len(ss) != 2 {
		return "", nil, false
	}

	az = ss[0]

	if strings.Index(ss[1], "v") != 0 {
		return "", nil, false
	}

	ss[1] = strings.Replace(ss[1], "v", "", 1)

	sss := strings.Split(ss[1], ".")
	if len(sss) != 2 {
		return "", nil, false
	}

	var err error

	sv = make([]int64, 2)

	sv[0], err = strconv.ParseInt(sss[0], 10, 64)
	if err != nil {
		return "", nil, false
	}

	sss[1] = strings.ReplaceAll(sss[1], "_", "")

	sv[1], err = strconv.ParseInt(sss[1], 10, 64)
	if err != nil {
		return "", nil, false
	}

	isRelease = true

	return az, sv, isRelease
}
