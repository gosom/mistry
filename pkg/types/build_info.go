package types

import (
	"fmt"
	"time"
)

// ContainerFailureExitCode is the exit code that signifies a failure
// before even running the container
const ContainerFailureExitCode = -999

// BuildInfo contains various information regarding the outcome of a
// particular build.
type BuildInfo struct {
	// Params are the job build parameters
	Params Params

	// Path is the absolute path where the build artifacts are located.
	Path string

	// Cached is true if the build artifacts were retrieved from the cache.
	Cached bool

	// Coalesced is true if the build was returned from another pending
	// build.
	Coalesced bool

	// Incremental is true if the results of a previous build were
	// used as the base for this build (ie. build cache).
	Incremental bool

	// ExitCode is the exit code of the container command.
	//
	// NOTE: irrelevant if Coalesced is true.
	ExitCode int

	// Err contains any errors that occured during the build.
	Err error

	// TransportMethod is the method with which the build artifacts can be
	// fetched.
	TransportMethod TransportMethod

	// StartedAt is the date and time when the build started.
	StartedAt time.Time

	// Log contains the stdout/stderr of the build.
	Log string

	// Errlog contains the stderr of the build.
	ErrLog string

	// URL is the relative URL at which the build log is available.
	URL string
}

type ErrImageBuild struct {
	Image string
	Err   error
}

func NewBuildInfo() *BuildInfo {
	bi := new(BuildInfo)
	bi.StartedAt = time.Now()
	bi.ExitCode = ContainerFailureExitCode

	return bi
}

func (e ErrImageBuild) Error() string {
	return fmt.Sprintf("could not build docker image '%s': %s", e.Image, e.Err)
}
