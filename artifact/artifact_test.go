//nolint:testpackage // Only same package cloud change variable privated
package artifact

import (
	"testing"
)

func TestArtifact_String(t *testing.T) {
	builtTime = "2026-02-16 12:00:00 +0000"
	commit = "5a966941"
	lastModifier = "John <john@example.com>"
	lastTag = "v1.0.0"

	art := Artifact{}
	got := art.String()

	want := `commit="5a966941" built="2026-02-16 12:00:00 +0000" tag="v1.0.0" modifier="John <john@example.com>"`

	if got != want {
		t.Fatalf("unexpected output:\n got: %s\n want: %s", got, want)
	}
}

func TestArtifact_BuiltTime(t *testing.T) {
	builtTime = "2026-02-16 12:00:00 +0000"

	art := Artifact{}
	got := art.BuiltTime()

	if got != builtTime {
		t.Fatalf("unexpected output:\n got: %s\n want: %s", got, builtTime)
	}
}

func TestArtifact_Commit(t *testing.T) {
	commit = "5a966941"

	art := Artifact{}
	got := art.Commit()

	if got != commit {
		t.Fatalf("unexpected output:\n got: %s\n want: %s", got, commit)
	}
}

func TestArtifact_LastModifier(t *testing.T) {
	lastModifier = "John <john@example.com>"

	art := Artifact{}
	got := art.LastModifier()

	if got != lastModifier {
		t.Fatalf("unexpected output:\n got: %s\n want: %s", got, lastModifier)
	}
}

func TestArtifact_LastTag(t *testing.T) {
	lastTag = "v1.0.0"

	art := Artifact{}
	got := art.LastTag()

	if got != lastTag {
		t.Fatalf("unexpected output:\n got: %s\n want: %s", got, lastTag)
	}
}
