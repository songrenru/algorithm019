/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// dp[i] = max(dp[i - 1], 0) + nums[i]
	l := len(nums)
	dp := nums[0]
	ans := dp
	for i := 1; i < l; i++ {
		dp = max(dp, 0) + nums[i]
		if dp > ans { ans = dp }
	}
	return ans
}

func max(i, j int) int {
	if i > j { return i }
	return j
}
// @lc code=end

