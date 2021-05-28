/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	// union find
	// dfs
	// bfs
	n := len(isConnected)
	visited := make([]bool, n)

	var dfs func(from int)
	dfs = func(from int) {
		visited[from] = true

		for to, isConn := range isConnected[from] {
			if isConn == 1 && !visited[to] {
				dfs(to)
			}
		}
	}

	count := 0
	for from, v := range visited {
		if !v {
			dfs(from)
			count++
		}
	}
	return count
}

// @lc code=end

