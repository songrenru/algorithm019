/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	// dp[i] = min(dp[i], dp[i - j^2] + 1)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		j := 1
		for true {
			tmp := i - j*j
			if tmp < 0 {
				break
			}
			if dp[i] > dp[tmp]+1 {
				dp[i] = dp[tmp] + 1
			}
			j++
		}
	}
	return dp[n]
}

// @lc code=end

