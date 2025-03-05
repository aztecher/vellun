package version

import (
	"fmt"
	"runtime"
)

// vellunVersion is set to Vellun's verson, revision and git author time reference during build
var vellunVersion string

// Version is the complete Vellun version string including Go version.
var Version string

func init() {
	Version = fmt.Sprintf(
		"%s go version %s %s/%s",
		vellunVersion,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
}
