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
		openCount int
	}
	layer1 := []variant{{"(", 1}}
	layer2 := []variant{}
	results := []string{}
	var closeCount int
	for viriantLen := 2; viriantLen < n*2; viriantLen++ {
		layer2 = []variant{}
		for _, l1v := range layer1 {
			closeCount = viriantLen - 1 - l1v.openCount
			if l1v.openCount < n {
				layer2 = append(
					layer2,
					variant{l1v.text + "(", l1v.openCount + 1},
				)
			}
			if closeCount < l1v.openCount {
				layer2 = append(
					layer2,
					variant{l1v.text + ")", l1v.openCount},
				)
			}
		}
		layer1 = layer2
	}
	for _, variant := range layer1 {
		results = append(results, variant.text+")")
	}
	return results
}