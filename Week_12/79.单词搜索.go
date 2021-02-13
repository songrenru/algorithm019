/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	// 定义方向
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
		// 越界检测
		if i < 0 || j < 0 || i == rows || j == cols {
			return false
		}
		// 不能重复
		c := board[i][j]
		if c == '#' {
			return false
		}
		// 单词前缀检测
		if c != word[k] {
			return false
		}
		if k == l - 1 {
			return true
		}
		// 继续遍历
		board[i][j] = '#'
		for d := 0; d < 4; d++ {
			if dfs(i + directs[d][0], j + directs[d][1], k + 1) {
				return true
			}
		}
		// 恢复
		board[i][j] = c
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

