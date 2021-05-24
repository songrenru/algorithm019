/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	// dp
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee) // 无行动/卖出
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	l := len(prices)
	hold := -prices[0]
	unHold := 0
	for i := 1; i < l; i++ {
		unHold, hold = max(unHold, hold+prices[i]-fee), max(hold, unHold-prices[i])
	}
	return unHold
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

