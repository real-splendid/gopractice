package leetcode

/*
36. Valid Sudoku
Determine if a 9 x 9 Sudoku board is valid.
Only the filled cells need to be validated according to the following rules:
Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

Example 1:
Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true

Example 2:
Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8.
Since there are two 8's in the top left 3x3 sub-box, it is invalid.
*/
func isValidSudoku(board [][]byte) bool {
	n := 9
	var hMap, vMap, bMap map[rune]bool
	var hKey, vKey, bKey rune
	var bI, bJ int
	for i := 0; i < n; i++ {
		hMap = map[rune]bool{}
		vMap = map[rune]bool{}
		bMap = map[rune]bool{}
		for j := 0; j < n; j++ {
			hKey = rune(board[i][j])
			vKey = rune(board[j][i])
			bI = 3*(i/3) + (j / 3)
			bJ = 3*(i%3) + (j % 3)
			bKey = rune(board[bI][bJ])
			if _, ok := hMap[hKey]; ok && hKey != '.' {
				return false
			}
			if _, ok := vMap[vKey]; ok && vKey != '.' {
				return false
			}
			if _, ok := bMap[bKey]; ok && bKey != '.' {
				return false
			}
			hMap[hKey], vMap[vKey], bMap[bKey] = true, true, true
		}
	}
	return true
}