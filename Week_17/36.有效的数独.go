/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	// 区域集合构造
	area1 := make([]bool, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			area[board[i][j]-'1'] = true
		}
	}
}

// @lc code=end
