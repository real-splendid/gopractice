package leetcode

/*
22. Generate Parentheses
Given n pairs of parentheses, write a function to generate all combinations
of well-formed parentheses.

Example 1:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:
Input: n = 1
Output: ["()"]
*/
func generateParenthesis(n int) []string {
	type variant struct {
		text      string
		len       int
		openCount int
	}
	resultVariantLen := n * 2
	buckets := map[int][]variant{1: {variant{"(", 1, 1}}}
	var results []string
	var closeCount int
	for size := 2; size < resultVariantLen; size++ {
		buckets[size] = []variant{}
		for _, v := range buckets[size-1] {
			closeCount = size - 1 - v.openCount
			if v.openCount < n {
				buckets[size] = append(
					buckets[size],
					variant{v.text + "(", size, v.openCount + 1},
				)
			}

			if closeCount < v.openCount {
				buckets[size] = append(
					buckets[size],
					variant{v.text + ")", size, v.openCount},
				)
			}
		}
	}
	for _, variant := range buckets[resultVariantLen-1] {
		results = append(results, variant.text+")")
	}
	return results
}