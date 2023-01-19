package leetcode

import (
	"reflect"
	"testing"
)

func TestCarFleet(t *testing.T) {
	check := func(want, got int, msg string) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("\n%s\nwant %v \ngot  %v", msg, want, got)
		}
	}

	check(
		3,
		carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}),
		"1",
	)
	check(
		1,
		carFleet(10, []int{3}, []int{3}),
		"2",
	)
	check(
		1,
		carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}),
		"3",
	)
	check(
		2,
		carFleet(10, []int{6, 8}, []int{3, 2}),
		"4",
	)
	check(
		4,
		carFleet(17, []int{8, 12, 16, 11, 7}, []int{6, 9, 10, 9, 7}),
		"5",
	)
}
