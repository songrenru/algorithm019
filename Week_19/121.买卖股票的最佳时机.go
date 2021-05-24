/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp[i][1] = min(dp[i-1][1], prices[i])
	// dp[i][0] = max(prices[i] - dp[i-1][1], dp[i-1][0])
	l := len(prices)
	dp := [2]int{}

	dp[0] = 0
	dp[1] = prices[0]
	for i := 1; i < l; i++ {
		dp[0], dp[1] = max(prices[i]-dp[1], dp[0]), min(dp[1], prices[i])
	}

	return dp[0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

