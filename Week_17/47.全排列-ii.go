/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	k := len(nums)
	res := [][]int{}
	visited := make([]bool, k)

	var dfs func(path []int)
	dfs = func(path []int) {
		l := len(path)
		if l == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i, num := range nums {
			if i > 0 && !visited[i-1] && nums[i] == nums[i-1] { // !visited[i-1]是关键
				continue
			}
			if visited[i] {
				continue
			}

			// backtracing
			visited[i] = true
			path = append(path, num)

			dfs(path)

			visited[i] = false
			path = path[:l]
		}
	}

	dfs([]int{})
	return res
}

// @lc code=end

