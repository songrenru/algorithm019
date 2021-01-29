/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp
	// dp + 空间优化（最值的自动匹配）
	// buy1[i] = max(buy1[i - 1] , -prices[i])
	// sell1[i] = max(buy[i - 1] + prices[i], sell1[i - 1])
	// buy2[i] = max(sell1[i - 1] - prices[i], buy2[i - 1])
	// sell2[i] = max(buy2[i - 1] + prices[i], sell2[i - 1])
	l := len(prices)
	if l <= 1 {
		return 0
	}
	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0
	for i := 1; i < l; i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(buy1+prices[i], sell1)
		buy2 = max(sell1-prices[i], buy2)
		sell2 = max(buy2+prices[i], sell2)
	}
	return sell2
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end
