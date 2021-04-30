/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp[i][j] = dp[i][j-1] + dp[i-1][j]
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if obstacleGrid[r][c] == 1 {
				dp[c] = 0
			} else if r == 0 {
				if c == 0 {
					dp[c] = 1
				} else {
					dp[c] = dp[c-1]
				}
			} else if c == 0 { // 不需要处理
				// if r == 0 {
				// 	dp[c] = 1
				// } else {
				// 	dp[c] = dp[c]
				// }
			} else {
				dp[c] = dp[c] + dp[c-1]
			}
		}
	}
	return dp[cols-1]
}

// @lc code=end

