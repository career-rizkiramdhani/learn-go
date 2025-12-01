package service

import (
	"fmt"
	"strings"
	"time"
)

// GenerateCode returns a unique code for a given module.
// It uses a module-specific prefix (if known), the current date (YYYYMMDD)
// and the current Unix time in milliseconds to ensure uniqueness.
// Example for module "user": GO_USR_20251201412512345
func GenerateCode(module string) string {
	m := strings.ToLower(strings.TrimSpace(module))

	prefixes := map[string]string{
		"user": "GO_USR_",
	}

	prefix, ok := prefixes[m]
	if !ok {
		// default prefix: GO_<UPPER_MODULE>_
		up := strings.ToUpper(m)
		if len(up) > 3 {
			up = up[:3]
		}
		prefix = fmt.Sprintf("GO_%s_", up)
	}

	date := time.Now().Format("20060102")
	// use milliseconds since epoch for good uniqueness while keeping length reasonable
	ts := time.Now().UnixNano() / int64(time.Millisecond)

	return fmt.Sprintf("%s%s%d", prefix, date, ts)
}

// GenerateUserCode is a convenience wrapper for the user module.
func GenerateUserCode() string {
	return GenerateCode("user")
}
