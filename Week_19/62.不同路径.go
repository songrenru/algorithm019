/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// 自顶向下，暴力递归，dp[i][j] = dp[i][j-1] + dp[i-1][j]
	// 自顶向下，带记忆的递归
	// 自底向上，dp // 进阶：空间复杂度优化
	if m > n {
		return uniquePaths(n, m)
	}

	dp := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[m-1]
}

// @lc code=end

