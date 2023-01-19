package leetcode

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	check := func(want, got []int) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("\nwant %v \ngot  %v", want, got)
		}
	}

	check(
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
		dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}),
	)

	check(
		[]int{1, 1, 1, 0},
		dailyTemperatures([]int{30, 40, 50, 60}),
	)

	check(
		[]int{1, 1, 0},
		dailyTemperatures([]int{30, 60, 90}),
	)
}
