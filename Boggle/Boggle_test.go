package boggle

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindWords(t *testing.T) {
	tests := map[string]struct {
		input [3][3]rune
		want  []string
	}{
		"simple": {input: [3][3]rune{{'G', 'I', 'Z'}, {'U', 'E', 'K'}, {'Q', 'S', 'E'}}, want: []string{"GEEKS", "QUIZ"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := FindWords(tc.input)

			diff := cmp.Diff(tc.want, got)

			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
