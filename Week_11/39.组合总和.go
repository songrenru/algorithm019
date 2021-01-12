/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	l := len(candidates)
	res := [][]int{}
	var dfs func(begin, sum int, path []int)
	dfs = func(begin, sum int, path []int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
	
		for i := begin; i < l; i++ {
			path = append(path, candidates[i])
			sum += candidates[i]

			dfs(i, sum, path)
			
			sum -= candidates[i]
			path = path[:len(path) - 1]
		}
	}
	dfs(0, 0, []int{})
	return res
}
// @lc code=end

