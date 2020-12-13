/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 暴力：求子序列，再求和，n^2
	// 分治法，NlogN
	// dp

	// f(j) = max(f(j - 1), 0) + nums[j]
	// dp[j] = max(d[j - 1], 0) + nums[j]
	// max
	res := nums[0]
	pre := nums[0] // dp只用存前一个子序的最大和
	lens := len(nums)
	for i := 1; i < lens; i++ {
		pre = max(pre, 0) + nums[i]
		if pre > res {
			res = pre
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
// @lc code=end

