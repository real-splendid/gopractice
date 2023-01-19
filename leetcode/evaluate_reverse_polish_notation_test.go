package leetcode

import (
	"reflect"
	"testing"
)

func TestEvaluateReversePolishNotation(t *testing.T) {
	check := func(want, got int) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("\nwant %v \ngot  %v", want, got)
		}
	}

	check(
		9,
		evalRPN([]string{"2", "1", "+", "3", "*"}),
	)
	check(
		6,
		evalRPN([]string{"4", "13", "5", "/", "+"}),
	)
	check(
		22,
		evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}),
	)
}
