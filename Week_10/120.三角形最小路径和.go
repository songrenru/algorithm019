/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	// 暴力递归
	// dp
	// 1. 递归 + 记忆化
	// 2. down to top
	// f(i, j) = min(f(i+1, j), f(i+1, j+1)) + triangle[i, j]
	dp := append([]int{}, triangle[len(triangle) - 1]...)
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func min(i, j int) int {
	if i < j { return i }
	return j
}
// @lc code=end

