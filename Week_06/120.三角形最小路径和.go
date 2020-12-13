/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	// f(i, j) = min(f(i+1, j), f(x+1, j+1)) + tri[i, j]
	// dp[i, j] = min(dp[i+1, j], dp[i+1, j+1]) + tri[i, j]
	r := len(triangle)
	dp := append([]int{}, triangle[r-1]...)
	for i := r - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
// @lc code=end

