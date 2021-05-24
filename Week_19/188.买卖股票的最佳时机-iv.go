/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	if k == 0 {
		return 0
	}

	l := len(prices)
	mid := l >> 1
	if k >= mid { // 允许交易次数，大于总次数，退化成贪心算法O(n)
		return greedy(prices)
	}

	// dp + 空间优化
	// dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i]) // j == 0,有特殊性 // 持股
	// dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + prices[i]) // 不持股
	dp := make([][2]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i][1] = -prices[0] // 第n次持股
		// dp[i][0] = 0 // 第n次未持股（未持股/卖出）
	}

	for i := 1; i < l; i++ {
		for j := 1; j <= k; j++ {
			// 无论题目中是否允许「在同一天买入并且卖出」这一操作，最终的答案都不会受到影响，这是因为这一操作带来的收益为零
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
		}
	}
	return dp[k][0]
}

func greedy(prices []int) int {
	l := len(prices)
	profit := 0
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

