/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	// dp
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]) // 不偷
	// dp[i][1] = dp[i-1][0] + nums[i] // 偷
	n := len(nums)
	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 0, nums[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]) // 不偷
		dp[i][1] = dp[i-1][0] + nums[i]        // 偷
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

