package leetcode

import (
// "sort"
// "strings"
)

/*
238. Product of Array Except Self
Given an integer array nums, return an array answer
such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/
func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	acc1 := 1
	acc2 := 1
	result := []int{}
	for range nums {
		result = append(result, 1)
	}
	for i := 0; i < numsLen; i++ {
		if i > 0 {
			acc1 *= nums[i-1]
		}
		result[i] = acc1
	}
	for i := numsLen - 1; i >= 0; i-- {
		if i < numsLen-1 {
			acc2 *= nums[i+1]
		}
		result[i] *= acc2
	}
	return result
}