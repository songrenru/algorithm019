/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	// dp 
	// dp[i] = min(dp[i], dp[i - j * j] + 1)
	dp := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i - j * j >= 0; j++ {
			dp[i] = min(dp[i], dp[i - j * j] + 1)
		}
	}
	return dp[n]
}

func min(i, j int) int {
	if i < j { return i }
	return j
}
// @lc code=end

