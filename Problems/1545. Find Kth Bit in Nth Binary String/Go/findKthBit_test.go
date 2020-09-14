package findKthBit

import (
	"fmt"
	"testing"
)

func TestFindKthBit(t *testing.T) {
	testCases := []struct {
		n, k int
		out byte
	}{
		{3, 1, '0'},
		{4, 11, '1'},
		{1, 1, '0'},
		{2, 3, '1'},
	}

	for _, tc := range testCases {
		tc := tc
		testname := fmt.Sprintf("FindKthBit(%d,%d)", tc.n, tc.k)
		t.Run(testname, func(t *testing.T) {
			r := FindKthBit(tc.n, tc.k)
			if r != tc.out {
				t.Errorf("'%c' = FindKthBit(%d, %d). Want '%c'", r, tc.n, tc.k, tc.out)
			}
		})
	}
}
