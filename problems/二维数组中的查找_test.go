package problems

import "testing"

func TestFindIn2DArray(t *testing.T) {
	tests := []struct {
		target int
		array  [][]int
		ok     bool
	}{
		{1, [][]int{{1, 2, 3}, {2, -10, 5}}, true},
		{345, [][]int{{2, -10, 5}}, false},
		{-345, [][]int{}, false},
	}

	for _, tt := range tests {
		ok := FindIn2DArray(tt.target, tt.array)
		if ok != tt.ok {
			t.Errorf("Test FindIn2DArray(%v, %v); got %v; expect %v",
				tt.target, tt.array, ok, tt.ok)
			t.Fail()
		}
	}
}
