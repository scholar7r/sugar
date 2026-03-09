package artifact_test

import (
	"testing"

	"github.com/scholar7r/sugar/artifact"
)

func TestArtifactString(t *testing.T) {
	//nolint:reassign // For test artifact BuildTime
	artifact.BuiltTime = "2026-02-16 12:00:00 +0000"
	//nolint:reassign // For test artifact Commit
	artifact.Commit = "abc123"
	//nolint:reassign // For test artifact LastModifier
	artifact.LastModifier = "John <john@example.com>"
	//nolint:reassign // For test artifact LastTag
	artifact.LastTag = "v1.0.0"

	a := artifact.Artifact{}
	got := a.String()

	want := `commit="abc123" built="2026-02-16 12:00:00 +0000" tag="v1.0.0" modifier="John <john@example.com>"`

	if got != want {
		t.Fatalf("unexpected output:\n got:  %s\n want: %s", got, want)
	}
}
