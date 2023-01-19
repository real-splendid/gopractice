package leetcode

import (
	"log"
	"strings"
)

/*
20. Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false
*/
func isValid(s string) bool {
	chars := strings.Split(s, "")
	pairs := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	log.Printf("\n%#v", chars)
	stack := []string{}
	stackLen := 0
	for _, c := range chars {
		if c == "{" || c == "[" || c == "(" {
			stack = append(stack, c)
			stackLen += 1
			continue
		}
		corr, ok := pairs[c]
		if ok && stackLen > 0 && stack[stackLen-1] == corr {
			stack = stack[:len(stack)-1]
			stackLen -= 1
			continue
		}
		return false
	}
	return len(stack) == 0
}

// func isValid2(s string) bool {
// 	// pairs := map[byte]byte{
// 	// 	'}': '{',
// 	// 	')': '(',
// 	// 	']': '[',
// 	// }
// 	stack := []byte{}
// 	lastByte := ' '
// 	for i := 0; i < len(s); i += 1 {
// 		switch s[i] {
// 		case '{', '[', '(':
// 			stack = append(stack, s[i])
// 		case '}', ']', ')':
// 			if len(stack) < 1 {
// 				return false
// 			}
// 			lastByte = stack[len(stack)-1]

// 			stack = stack[:len(stack)-1]
// 		default:
// 			return false
// 		}
// 	}
// 	return len(stack) == 0
// }
