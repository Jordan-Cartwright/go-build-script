package version

import (
	"fmt"
)

var (
	// Deliberately uninitialized, see GetVersion()
	version   string
	gitCommit string
)

func GetFullVersion() string {
	return fmt.Sprintf("%s-%s", GetVersion(), GetCommit())
}

func GetCommit() string {
	if gitCommit != "" {
		return gitCommit
	}
	return "unknown"
}

func GetVersion() string {
	if version != "" {
		return version
	}
	return "unknown"
}
