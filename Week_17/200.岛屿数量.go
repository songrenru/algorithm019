/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	// dfs
	// union-find disjoint sets
	// bfs
	r, c := len(grid), len(grid[0])
	num := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				num++
				bfs(i, j, r, c, grid)
			}
		}
	}
	return num
}

var directs = [4][2]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

func bfs(i, j, r, c int, grid [][]byte) {
	q := []int{i, j}
	for len(q) > 0 {
		size := len(q)
		for k := 0; k < size; k += 2 {
			x, y := q[k], q[k+1]
			grid[x][y] = '0'

			for _, direct := range directs {
				x1, y1 := x+direct[0], y+direct[1]
				if x1 < 0 || y1 < 0 || x1 == r || y1 == c {
					continue
				}
				if grid[x1][y1] == '1' {
					q = append(q, x1, y1)
				}
			}
		}
		q = q[size:]
	}
}

// @lc code=end

