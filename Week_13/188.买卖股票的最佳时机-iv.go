import "math"

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	// dp
	// dp[i][j][2] // i == len(prices),j in [0, k)
	// dp[i][j][1] = max(dp[i - 1][j][1], dp[i - 1][j][0] - prices[i]) // j == 0,有特殊性 // 持股
	// dp[i][j][0] = max(dp[i - 1][j][0], dp[i - 1][j - 1][1] + prices[i]) // 不持股
	l := len(prices)
	if k == 0 || l <= 1 {
		return 0
	}
	if k >= l/2 {
		return greedy(prices)
	}
	dp := make([][2]int, k+1)
	// dp[0][1] = -prices[0]
	// dp[0][0] = 0
	for i := 0; i <= k; i++ {
		dp[i][1] = math.MinInt64
		// dp[i][0] = 0
	}
	for i := 0; i < l; i++ {
		for j := 1; j <= k; j++ {
			tmp := dp[j][1]
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i]) // 持股
			dp[j][0] = max(dp[j][0], tmp+prices[i])        // 不持股
		}
	}
	return dp[k][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func greedy(prices []int) int {
	profit := 0
	l := len(prices)
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// @lc code=end
