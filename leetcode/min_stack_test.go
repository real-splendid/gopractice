package leetcode

import (
	"testing"
	// "reflect"
	"log"
)

/*
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
*/
func TestMinStack(t *testing.T) {
	minStack := Constructor()
	log.Printf("\n%#v\n", minStack)
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	want := -3
	got := minStack.GetMin()
	if want != got {
		t.Errorf("\nwant %v \ngot  %v", want, got)
	}

	minStack.Pop()
	want = 0
	got = minStack.Top()
	if want != got {
		t.Errorf("\nwant %v \ngot  %v", want, got)
	}

	want = -2
	got = minStack.GetMin()
	if want != got {
		t.Errorf("\nwant %v \ngot  %v", want, got)
	}
}
