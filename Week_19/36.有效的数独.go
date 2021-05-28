/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	// 三次遍历
	// 一次遍历
	// 一次遍历 + 位运算
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxes := [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			intVal := board[i][j] - '1'
			// 行判定
			if rows[i][intVal] {
				return false
			}

			// 列判定
			if cols[j][intVal] {
				return false
			}

			// 宫格判定
			boxIdx := (i/3)*3 + (j / 3)
			if boxes[boxIdx][intVal] {
				return false
			}

			rows[i][intVal] = true
			cols[j][intVal] = true
			boxes[boxIdx][intVal] = true
		}
	}

	return true
}

// @lc code=end

