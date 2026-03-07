package strung_test

import (
	"slices"
	"testing"

	"github.com/scholar7r/sugar/strung"
)

func TestTrimAround(t *testing.T) {
	tests := []struct {
		name   string
		v      []string
		cutset string
		want   []string
	}{
		{
			name:   "trim spaces around",
			v:      []string{" ADD ", " 1 ", " AND ", " 2 "},
			cutset: " ",
			want:   []string{"ADD", "1", "AND", "2"},
		},
		{
			name:   "trim dots around",
			v:      []string{".ADD.", ".1.", ".AND.", ".2."},
			cutset: ".",
			want:   []string{"ADD", "1", "AND", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strung.TrimAround(tt.v, tt.cutset)
			if slices.Compare(got, tt.want) != 0 {
				t.Errorf("TrimAround() = %v, want %v", got, tt.want)
			}
		})
	}
}
