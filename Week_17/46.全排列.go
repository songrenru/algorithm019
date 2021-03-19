/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	// backtracing
	visited := map[int]bool{}
	res := [][]int{}
	k := len(nums)

	var dfs func(path []int)
	dfs = func(path []int) {
		l := len(path)
		if l == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for _, num := range nums {
			if !visited[num] {
				visited[num] = true
				path = append(path, num)

				dfs(path)

				visited[num] = false
				path = path[:l]
			}
		}
	}

	dfs([]int{})
	return res
}

// @lc code=end

