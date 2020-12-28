/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	directs := [4][2]int{
		{-1, 0}, // up
		{1, 0}, // down
		{0, -1}, // left
		{0, 1}, // right
	}
	rows, cols := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i == rows || j == cols {
			return
		}
	
		if grid[i][j] == '0' {
			return
		}
	
		grid[i][j] = '0'
		
		for k := 0; k < 4; k++ {
			dfs(i + directs[k][0], j + directs[k][1])
		}
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}
	return count
}


// @lc code=end

