/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp + 空间优化
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 未持股（卖出/无动作）
	// dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	l := len(prices)
	if l <= 1 {
		return 0
	}
	// dp := make([][2]int, l)
	// // dp[0][0] = 0
	// dp[0][1] = -prices[0]
	// dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
	// dp[1][1] = max(dp[0][1], -prices[1])
	prevProfit0 := 0
	profit0 := 0
	profit1 := -prices[0]
	for i := 1; i < l; i++ {
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 未持股（卖出/无动作）
		// dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		newProfit0 := max(profit0, profit1+prices[i])
		newProfit1 := max(profit1, prevProfit0-prices[i])
		prevProfit0 = profit0
		profit0 = newProfit0
		profit1 = newProfit1
	}
	// return dp[l-1][0]
	return profit0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

