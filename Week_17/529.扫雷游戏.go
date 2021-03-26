/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	directs := [8][2]int{
		{-1, 0},  // 上
		{-1, 1},  // 右上
		{0, 1},   // 右
		{1, 1},   // 右下
		{1, 0},   // 下
		{1, -1},  // 左下
		{0, -1},  // 左
		{-1, -1}, // 左上
	}

	r := len(board)
	c := len(board[0])

	start := [2]int{click[0], click[1]}
	visited := map[[2]int]bool{start: true}
	q := [][2]int{start}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			x, y := q[i][0], q[i][1]

			count := 0
			set := [][2]int{}
			for j := 0; j < 8; j++ {
				newX := x + directs[j][0]
				newY := y + directs[j][1]

				if newX < 0 || newY < 0 || newX == r || newY == c {
					continue
				}
				if board[newX][newY] == 'M' {
					count++
				}

				if count == 0 {
					tmp := [2]int{newX, newY}
					if !visited[tmp] {
						set = append(set, tmp)
					}
				}
			}

			if count == 0 {
				board[x][y] = 'B'
				q = append(q, set...)
				for _, v := range set {
					visited[v] = true
				}
			} else {
				board[x][y] = byte(count + 48)
			}
		}
		q = q[size:]
	}
	return board
}

// @lc code=end

