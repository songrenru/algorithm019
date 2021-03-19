/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	// dfs（有或没有的穷举） + 剪枝
	// backtracing(for循环)
	res := [][]int{}
	var dfs func(i int, path []int)
	dfs = func(i int, path []int) {
		// recursion terminator
		l := len(path)
		if l == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		// 剪枝
		if n-i < k-l-1 {
			return
		}

		for j := i; j <= n; j++ {
			path = append(path, j)

			dfs(j+1, path)

			path = path[:l]
		}
	}

	dfs(1, []int{})
	return res
}

// @lc code=end

