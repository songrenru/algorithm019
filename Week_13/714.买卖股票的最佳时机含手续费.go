/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	// dp
	// dp + 空间优化
	// dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] - prices[i]) // 持股
	// dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] + prices[i] - fee) // 不持股
	l := len(prices)
	if l <= 1 {
		return 0
	}
	hold := -prices[0]
	unhold := 0
	for i := 1; i < l; i++ {
		hold, unhold = max(hold, unhold-prices[i]), max(unhold, hold+prices[i]-fee)
	}
	return unhold
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end
