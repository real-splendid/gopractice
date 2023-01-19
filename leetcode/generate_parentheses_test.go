package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	check := func(want, got []string) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("\nwant %v \ngot  %v", want, got)
		}
	}

	check(
		[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		generateParenthesis(3),
	)
	check(
		[]string{"()"},
		generateParenthesis(1),
	)
}
