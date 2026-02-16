// Package artifact exposes build-time metadata injected via -ldflags.
//
// It provides information such as the Git commit hash, build timestamp,
// last Git tag, and last commit author. These values are typically
// populated during the build process using:
//
//	go build -ldflags "-X <module>/artifact.Commit=..."
package artifact

import "fmt"

type Artifact struct{}

var (
	BuiltTime    string // date +'%F %T %z'
	Commit       string // git rev-parse --short HEAD
	LastModifier string // git show -s --format='format:%aN <%ae>' HEAD
	LastTag      string // git describe --tags --abbrev=0 2>/dev/null || echo "none"
)

func (Artifact) String() string {
	return fmt.Sprintf("commit=\"%s\" built=\"%s\" tag=\"%s\" modifier=\"%s\"",
		Commit, BuiltTime, LastTag, LastModifier)
}
