package leetcode

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	check := func(want, got []int) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("\nwant %v \ngot  %v", want, got)
		}
	}

	check(
		[]int{24, 12, 8, 6},
		productExceptSelf([]int{1, 2, 3, 4}),
	)
	check(
		[]int{0, 0, 9, 0, 0},
		productExceptSelf([]int{-1, 1, 0, -3, 3}),
	)
}
