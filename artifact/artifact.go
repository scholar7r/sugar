// Package artifact exposes build-time metadata injected via -ldflags.
//
// It provides information such as the Git commit hash, build timestamp,
// last Git tag, and last commit author. These values are typically
// populated during the build process using:
//
//	go build -ldflags "-X <module>/artifact.Commit=..."
//
// The Print function outputs the metadata in a single-line, log-friendly format.
package artifact

import "fmt"

var (
	BuiltTime    string // date +'%F %T %z'
	Commit       string // git rev-parse --short HEAD
	LastModifier string // git show -s --format='format:%aN <%ae>' HEAD
	LastTag      string // git describe --tags --abbrev=0 2>/dev/null || echo "none"
)

// Print writes the build metadata to standard output in a single-line,
// log-friendly format. The output includes commit hash, build time,
// last tag, and last modifier information.
func Print() {
	fmt.Printf("artifact commit=\"%s\" built=\"%s\" tag=\"%s\" modifier=\"%s\"\n",
		Commit, BuiltTime, LastTag, LastModifier)
}
