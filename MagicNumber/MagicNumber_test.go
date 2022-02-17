package magicnumber

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNthMagicNumber(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"2": {input: 2, want: 25},
		"5": {input: 5, want: 130},
		"6": {input: 6, want: 150},
		"7": {input: 7, want: 155},
		"8": {input: 8, want: 625},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := NthMagicNumber(tc.input)

			diff := cmp.Diff(tc.want, got)

			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
