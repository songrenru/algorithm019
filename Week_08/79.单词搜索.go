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

	var dfs func(i, j, deep int, path []byte) bool
	dfs = func(i, j, deep int, path []byte) bool {
		// recursion terminator
		if i < 0 || j < 0 || i == rows || j == cols {
			return false
		}
		if board[i][j] == '1' {
			return false
		}
		if word[deep] != board[i][j] {
			return false
		}
		// process current logic
		// drill down
		oldChar := board[i][j]
		board[i][j] = '1'
		path = append(path, oldChar)
		deep++
		if deep == l && string(path) == word {
			return true
		}
		
		for k := 0; k < 4; k++ {
			if dfs(i + directs[k][0], j + directs[k][1], deep, path) {
				return true
			}
		}
		// revert status if needed
		path = path[:len(path) - 1]
		board[i][j] = oldChar
		return false
	}
	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0, []byte{}) {
				return true
			}
		}
	}
	return false
}
// @lc code=end

