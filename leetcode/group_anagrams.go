package leetcode

import (
	"sort"
	"strings"
)

/*
49. Group Anagrams
Given an array of strings strs, group the anagrams together.
You can return the answer in any order.

Example:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/
func groupAnagrams(strs []string) [][]string {
	keys := map[string]int{}
	results := [][]string{}
	resultsLen := 0

	for _, w := range strs {
		strSlice := strings.Split(w, "")
		sort.Strings(strSlice)
		key := strings.Join(strSlice, "")
		keyR, ok := keys[key]
		if !ok {
			keys[key] = resultsLen
			keyR = resultsLen
			results = append(results, []string{})
			resultsLen = len(results)
		}
		results[keyR] = append(results[keyR], w)
	}
	return results
}
