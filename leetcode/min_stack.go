package leetcode

import (
// "log"
// "math"
)

/*
155. Min Stack
Design a stack that supports push, pop, top,
and retrieving the minimum element in constant time.

Implement the MinStack class:
MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.

Example 1:
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
type MinStack struct {
	data []int
	mins []int
	len  int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, 0}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	min := val
	prevMin := this.GetMin()
	if prevMin < min {
		min = prevMin
	}
	this.mins = append(this.mins, min)
	this.len = this.len + 1
}

func (this *MinStack) Pop() {
	this.data = this.data[0 : this.len-1]
	this.mins = this.mins[0 : this.len-1]
	this.len = this.len - 1
}

func (this *MinStack) Top() int {
	return this.data[this.len-1]
}

func (this *MinStack) GetMin() int {
	if this.len == 0 {
		return 9223372036854775807
	}
	return this.mins[this.len-1]
}