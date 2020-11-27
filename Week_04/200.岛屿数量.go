/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	var dfs func(r, c int)
	dfs = func(r, c int) {
		// recursion terminator
		if !inArea(grid, r, c) || grid[r][c] != '1' {
			return
		}
		// process current logic
		grid[r][c] = 2
		// drill down
		dfs(r - 1, c);
		dfs(r + 1, c);
		dfs(r, c - 1);
		dfs(r, c + 1);
	}

	rowNum := len(grid)
	colNum := len(grid[0])
	count := 0
	for r := 0; r < rowNum; r++ {
		for c := 0; c < colNum; c++ {
			if grid[r][c] == '1' {
				dfs(r, c)
				count++
			}
		}
	}
	return count
}

func inArea(grid [][]byte, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
// @lc code=end

