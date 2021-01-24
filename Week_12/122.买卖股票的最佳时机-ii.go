/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// greedy
	// dp
	l := len(prices)
	if l <= 1 {
		return 0
	}

	dp := make([][2]int, l)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
		dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])
	}
	return dp[l - 1][0]
}

func max(i, j int) int {
	if i > j { return i }
	return j
}
// @lc code=end

