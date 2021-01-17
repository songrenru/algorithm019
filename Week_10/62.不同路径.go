/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// dp[i][j] = dp[i][j - 1] + dp[i - 1][j]
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] = dp[j] + dp[j - 1]
			}
		}
	}
	return dp[n - 1]
}
// @lc code=end

