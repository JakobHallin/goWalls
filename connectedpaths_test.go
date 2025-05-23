package paths

import (
	"testing"
)

func TestPathEnumeration(t *testing.T) {
	cases := []struct{ cols, rows, paths int }{
		{0, 0, 0},
		{1, 1, 1},
		{1, 2, 2},
		{2, 1, 1},
		{2, 2, 4},
		{3, 3, 17},
		{5, 3, 99},
		{5, 4, 178},
		{5, 5, 259},
		{6, 4, 466},
	}
	for _, test := range cases {
		var res [][]int
		res = ConnectedPaths(test.cols, test.rows)
		if len(res) != test.paths {
			t.Errorf("wrong number of paths for grid %dx%d, expected %d, got %d",
				test.cols, test.rows, test.paths, len(res))
		}
	}
}
