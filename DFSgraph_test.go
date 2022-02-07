package dfs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestDFS(t *testing.T) {
	tests := map[string]struct {
		input [][2]int
		want  []int
	}{
		"simple":       {input: [][2]int{[2]int{0, 1}, [2]int{0, 2}, [2]int{1, 2}, [2]int{2, 0}, [2]int{2, 3}, [2]int{3, 3}}, want: []int{0, 1, 2, 3}},
		"disconnected": {input: [][2]int{[2]int{0, 1}, [2]int{0, 2}, [2]int{1, 2}, [2]int{2, 0}, [2]int{3, 3}}, want: []int{0, 1, 2, 3}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			g := new(Graph)
			g.AddNode()
			g.AddNode()
			g.AddNode()
			g.AddNode()
			for _, edge := range tc.input {
				g.AddEdge(edge[0], edge[1], 1)
			}
			got := g.DFS()
			less := func(a, b int) bool { return a < b }

			diff := cmp.Diff(tc.want, got, cmpopts.SortSlices(less))

			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
