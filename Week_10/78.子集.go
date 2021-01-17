/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	l := len(nums)
	res := [][]int{}

	var dfs func(i int, path []int)
	dfs = func(i int, path []int) {
		if i == l {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
	
		dfs(i + 1, path)
	
		path = append(path, nums[i])
		dfs(i + 1, path)
	}
	dfs(0, []int{})
	return res
}

// @lc code=end

