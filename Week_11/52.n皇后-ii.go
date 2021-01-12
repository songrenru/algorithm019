/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start
func totalNQueens(n int) int {
	col := map[int]bool{}
	pie := map[int]bool{}
	na := map[int]bool{}

	count := 0
	var dfs func(r int, path []int)
	dfs = func(r int, path []int) {
		if r == n {
			count++
			return 
		}
	
		for c := 0; c < n; c++ {
			if col[c] || pie[r + c] || na[r - c] {
				continue
			}
			path = append(path, c)
			col[c] = true
			pie[r + c] = true
			na[r - c] = true
	
			dfs(r + 1, path)
			// backtrace，为下一层循环做准备
			path = path[:len(path) - 1] 
			col[c] = false
			pie[r + c] = false
			na[r - c] = false
		}
	}
	dfs(0, []int{})
	return count
}
// @lc code=end

