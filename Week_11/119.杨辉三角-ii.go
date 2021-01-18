/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	// dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j]
	// 从后往前，不需要pre变量
	dp := make([]int, rowIndex + 1)
	for i := 0; i <= rowIndex; i++ {
		// pre := 1
		for j := i; j >= 0; j-- {
			if j == 0 || j == i {
				dp[j] = 1
			} else {
				// temp := dp[j]
				// dp[j] = pre + dp[j]
				// pre = temp
				dp[j] = dp[j - 1] + dp[j]
			}
		}
	}
	return dp
}
// @lc code=end
