/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}
	lens := len(nums)

	var helper func(idx int, temp []int)
	helper = func(idx int, temp []int)  {
		// recursion terminator
		if idx == lens {
			res = append(res, temp)
			return
		}
		// process current logic
		// drill down
		helper(idx + 1, temp)
		temp = append(temp, nums[idx])
		helper(idx + 1, temp)
		// reverse state
	}
	helper(0, nil)
	return res
}
// @lc code=end

