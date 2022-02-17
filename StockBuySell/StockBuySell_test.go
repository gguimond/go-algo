package stockbuysell

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNthMagicNumber(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"nominal":   {input: []int{100, 180, 260, 310, 40, 535, 695}, want: 865},
		"no profit": {input: []int{100, 50, 10}, want: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := StockBuySell(tc.input)

			diff := cmp.Diff(tc.want, got)

			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestNthMagicNumberPanic(t *testing.T) {
	t.Run("not enough prices", func(t *testing.T) {
		defer func() { recover() }()
		StockBuySell([]int{100})
		t.Fatalf("did not panic")
	})
}
