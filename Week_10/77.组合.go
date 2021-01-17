/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}

	var dfs func(i int, path []int)
	dfs = func(i int, path []int) {
		// 剪枝
		if n - i + 1 < k - len(path) {
			return
		}

		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
	
		for j := i; j <= n; j++ {
			path = append(path, j)
	
			dfs(j + 1, path)
	
			path = path[:len(path) - 1]
		}
	}

	dfs(1, []int{})
	return res

}


// @lc code=end

