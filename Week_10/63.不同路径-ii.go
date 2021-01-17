/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])

	dp := make([]int, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				if i == 0 && j == 0 {
					dp[j] = 1
				} else if i == 0 {
					dp[j] = dp[j - 1]
				} else if j == 0 {
					continue
				} else {
					dp[j] = dp[j] + dp[j - 1]
				}
			}
		}
	}
	return dp[cols - 1]
}
// @lc code=end

