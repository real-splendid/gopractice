package leetcode

import (
	"sort"
)

/*
347. Top K Frequent Elements
Given an integer array nums and an integer n, return the n most frequent elements.
You may return the answer in any order.

Example 1:
Input: nums = [1,1,1,2,2,3], n = 2
Output: [1,2]

Example 2:
Input: nums = [1], n = 1
Output: [1]

Follow up: Your algorithm's time complexity must be better than O(n log n),
where n is the array's size.
*/
func topKFrequent(nums []int, n int) []int {
	keys := []int{}
	freq := map[int]int{}
	for _, v := range nums {
		freq[v] += 1
	}
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})
	return keys[:n]
}

func topKFrequentV2(nums []int, n int) []int {
	freq := map[int]int{}
	buckets := make([][]int, len(nums)+1)
	result := []int{}
	for _, v := range nums {
		freq[v] += 1
	}
	for k, v := range freq {
		buckets[v] = append(buckets[v], k)
	}
	for i := len(nums); i > 0; i -= 1 {
		for j := range buckets[i] {
			result = append(result, buckets[i][j])
			if len(result) == n {
				return result
			}
		}
	}
	return result
}
