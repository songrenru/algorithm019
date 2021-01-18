/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	// dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j]
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i + 1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = dp[i - 1][j - 1] + dp[i - 1][j]
			}
		}
		dp[i] = row
	}
	return dp
}
// @lc code=end

