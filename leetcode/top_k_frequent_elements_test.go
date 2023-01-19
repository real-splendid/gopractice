package leetcode

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	check := func(want, got []int) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("\nwant %v \ngot  %v", want, got)
		}
	}

	check(
		[]int{1, 2},
		topKFrequentV2([]int{3, 1, 1, 1, 2, 2}, 2),
	)
	check(
		[]int{1},
		topKFrequentV2([]int{1}, 1),
	)
	check(
		[]int{4444, 333},
		topKFrequentV2([]int{1, 22, 333, 333, 4444, 4444, 4444}, 2),
	)
}
