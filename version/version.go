package version

import (
	"fmt"
	"runtime"
	"strings"
)

// AppRevision ...
const AppRevision = "dev"

// AppRevisionTag ...
const AppRevisionTag = ""

// AppRevisionOrTag ...
func AppRevisionOrTag() string {
	if AppRevisionTag == "" {
		return AppRevision
	}
	return AppRevisionTag
}

// AppRevisionWithTag ...
func AppRevisionWithTag() string {
	return AppRevision + ":" + AppRevisionTag
}

// UserAgent ...
func UserAgent() string {
	var b strings.Builder
	defer b.Reset()

	fmt.Fprintf(&b, "kong-client/%s ", AppRevisionOrTag())
	fmt.Fprintf(&b, "(%s; %s) ", runtime.GOOS, runtime.GOARCH)
	fmt.Fprintf(&b, "%s", runtime.Version())

	s := b.String()
	return s
}
