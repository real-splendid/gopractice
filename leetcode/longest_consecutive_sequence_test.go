package leetcode

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	check := func(want, got int, msg string) {
		if want != got {
			t.Errorf("\nfailed test %v\nwant %v \ngot  %v", msg, want, got)
		}
	}
	check(
		4,
		longestConsecutive([]int{100, 4, 200, 1, 3, 2}),
		"test 1",
	)
	check(
		9,
		longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}),
		"test 2",
	)
	check(
		1,
		longestConsecutive([]int{0}),
		"test 3",
	)
	check(
		1,
		longestConsecutive([]int{0, 0}),
		"test 4",
	)
	check(
		7,
		longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}),
		"test 5",
	)
}