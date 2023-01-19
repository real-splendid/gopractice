package leetcode

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	check := func(want, got bool, msg string) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("%v\nwant %v \ngot  %v", msg, want, got)
		}
	}

	check(
		true,
		isValid("()"),
		"1",
	)
	check(
		true,
		isValid("()[]{}"),
		"2",
	)
	check(
		false,
		isValid("(]"),
		"3",
	)
	check(
		false,
		isValid("]"),
		"4",
	)
}
