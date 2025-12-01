package service

import (
	"fmt"
	"strings"
	"time"
)

// GenerateCode returns a unique code for a given module.
// Caller may provide a `prefix` (e.g. "GO_USR_"). If `prefix` is empty,
// a default prefix `GO_<UPPERSHORT>_` will be used where <UPPERSHORT>
// is the module name uppercased and truncated to 3 chars.
// The generated format is: <PREFIX><YYYYMMDD><unix-milliseconds>
// Example: GO_USR_20251201<ms>
func GenerateCode(module, prefix string) string {
	m := strings.ToLower(strings.TrimSpace(module))

	if strings.TrimSpace(prefix) == "" {
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

// GenerateUserCode is a convenience wrapper for the user module that accepts a prefix.
func GenerateUserCode(prefix string) string {
	return GenerateCode("user", prefix)
}
