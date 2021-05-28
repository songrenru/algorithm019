/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	// dfs + backtrace

	// 1. 构建数据
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxes := [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			intVal := board[i][j] - '1'
			boxIdx := (i/3)*3 + (j / 3)

			rows[i][intVal] = true
			cols[j][intVal] = true
			boxes[boxIdx][intVal] = true
		}
	}

	// 2. dfs + 回溯
	var dfs func(site int) bool
	dfs = func(site int) bool {
		if site == 81 {
			return true
		}

		row := site / 9
		col := site % 9

		if board[row][col] != '.' {
			return dfs(site + 1)
		}

		boxIdx := (row/3)*3 + (col / 3)
		for num := 0; num < 9; num++ {
			// 有效性判定
			if rows[row][num] || cols[col][num] || boxes[boxIdx][num] {
				continue
			}

			byteVal := (byte)('1' + num)
			// 回溯
			rows[row][num] = true
			cols[col][num] = true
			boxes[boxIdx][num] = true
			board[row][col] = byteVal

			if dfs(num + 1) {
				return true
			}

			rows[row][num] = false
			cols[col][num] = false
			boxes[boxIdx][num] = false
			board[row][col] = '.'
		}
		return false
	}

	dfs(0)
}

// @lc code=end

