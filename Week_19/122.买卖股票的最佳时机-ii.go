/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// greedy
	// dp
	// dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
	// dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	l := len(prices)

	profit := 0
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// @lc code=end

