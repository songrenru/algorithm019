/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	// 暴力
	// dp
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	// dp[i][1] = dp[i-1][0] + nums[i]
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	dp := make([][2]int, lens)
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < lens; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[lens-1][0], dp[lens-1][1])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
// @lc code=end

