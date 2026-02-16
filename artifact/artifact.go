// Package artifact exposes build-time metadata injected via -ldflags.
//
// It provides information such as the Git commit hash, build timestamp,
// last Git tag, and last commit author. These values are typically
// populated during the build process using:
//
//	go build -ldflags "-X <module>/artifact.Commit=..."
package artifact

import "fmt"

// Artifact represents build metadata of the current binary.
//
// It is a zero-sized type used to provide a String method
// for formatted output of build information.
type Artifact struct{}

var (
	// BuiltTime is the build timestamp.
	// Example: `date +'%F %T %z'`
	BuiltTime string

	// Commit is the short Git commit hash.
	// Example: `git rev-parse --short HEAD`
	Commit string

	// LastModifier is the author of the latest commit.
	// Example: `git show -s --format='format:%aN <%ae>' HEAD`
	LastModifier string

	// LastTag is the most recent Git tag.
	// Example:
	//   git describe --tags --abbrev=0 2>/dev/null || echo "none"
	LastTag string
)

// String returns a formatted representation of the build metadata.
func (Artifact) String() string {
	return fmt.Sprintf(`commit="%s" built="%s" tag="%s" modifier="%s"`,
		Commit, BuiltTime, LastTag, LastModifier)
}
