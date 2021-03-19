/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	col := make([]bool, n)
	pie := make([]bool, 2*n-1)
	na := map[int]bool{}
	res := [][]string{}
	var backtracing func(r int, path []int)
	backtracing = func(r int, path []int) {
		if r == n {
			res = append(res, generateBoard(path, n))
			return
		}

		for c := 0; c < n; c++ {
			if col[c] || pie[r+c] || na[r-c] {
				continue
			}
			col[c], pie[r+c], na[r-c] = true, true, true
			path = append(path, c)

			backtracing(r+1, path)

			path = path[:len(path)-1]
			col[c], pie[r+c], na[r-c] = false, false, false
		}
	}

	backtracing(0, []int{})
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

