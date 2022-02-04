package finder_test

import (
	"fmt"
	"math/rand"
	"testing"
	"atomsbt.20200201/task2"
)

func TestFindElement(t *testing.T) {
	tests := []struct {
		in      []int
		want    int
		wantErr error
	}{
		{createInput(2), 2, nil},
		{createInput(100), 100, nil},
		{[]int{1, 2}, 0, finder.ErrMalformedInput},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got, gotErr := finder.FindElement(test.in)
			if gotErr != test.wantErr {
				t.Fatalf("Got %v error, want %v", gotErr, test.wantErr)
			}
			if got != test.want {
				t.Fatalf("Got %d value, want %d", got, test.want)
			}
		})
	}
}

// returns slice of ints:
// elements: 1 through 100, n-th element is replaced with 0;
// slice then shuffled
func createInput(n int) []int {
	res := make([]int, 100)
	for i := 0; i < 100; i++ {
		if i == n-1 {
			res[i] = 0
		} else {
			res[i] = i + 1
		}

	}
	perm(res)
	return res
}
func perm(is []int) {
	for i := range is {
		j := rand.Intn(i + 1)
		is[i], is[j] = is[j], is[i]
	}
}
