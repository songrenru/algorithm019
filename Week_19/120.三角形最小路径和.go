/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	// 自顶向下，暴力递归
	// dp
	// 1. 递归 + 记忆化
	// 2. down to top

	// dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + triangle[i][j]
	rows := len(triangle)
	// cols := len(triangle[rows-1])

	dp := append([]int{}, triangle[rows-1]...)
	for r := rows - 2; r >= 0; r-- { // 倒数第二行开始
		l := len(triangle[r])
		for c := 0; c < l; c++ {
			dp[c] = min(dp[c], dp[c+1]) + triangle[r][c]
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

