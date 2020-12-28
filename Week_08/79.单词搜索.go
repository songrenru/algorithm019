/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	directs := [4][2]int{
		{-1, 0}, // up
		{1, 0}, // down
		{0, -1}, // left
		{0, 1}, // right
	}
	rows, cols := len(board), len(board[0])
	l := len(word)

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// recursion terminator
		// 边界检测
		if i < 0 || j < 0 || i == rows || j == cols {
			return false
		}
		// 去重
		if board[i][j] == '1' {
			return false
		}
		// 剪枝
		if board[i][j] != word[k] {
			return false
		}
		// 找到字符串
		if k == l - 1 {
			return true
		}
		// process current logic
		old := board[i][j]
		board[i][j] = '1'
		// dirll down
		for m := 0; m < 4; m++ {
			if dfs(i + directs[m][0], j + directs[m][1], k + 1) {
				return true
			}
		}
		// revert status if needed
		board[i][j] = old
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}


// @lc code=end

