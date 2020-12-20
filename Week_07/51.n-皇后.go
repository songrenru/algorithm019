/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	col := map[int]bool{}
	pie := map[int]bool{}
	na := map[int]bool{}

	res := [][]string{}
	var dfs func(r int, path []int)
	dfs = func(r int, path []int) {
		if r == n {
			res = append(res, generateBoard(path, n))
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
	return res
}

func generateBoard(queues []int, n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		temp := make([]byte, n)
		for j := 0; j < n; j++ {
			if j == queues[i] {
				temp[j] = 'Q'
			} else {
				temp[j] = '.'
			}
		}
		res[i] = string(temp)
	}
	return res
}
// @lc code=end

