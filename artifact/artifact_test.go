package artifact

import "testing"

func TestArtifactString(t *testing.T) {
	BuiltTime = "2026-02-16 12:00:00 +0000"
	Commit = "abc123"
	LastModifier = "John <john@example.com>"
	LastTag = "v1.0.0"

	a := Artifact{}
	got := a.String()

	want := `commit="abc123" built="2026-02-16 12:00:00 +0000" tag="v1.0.0" modifier="John <john@example.com>"`

	if got != want {
		t.Fatalf("unexpected output:\n got:  %s\n want: %s", got, want)
	}
}
