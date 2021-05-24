/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp
	// buy1[i] = max(buy1[i-1], -priecs[i]) // 一次交易，买入
	// sell1[i] = max(sell1[i-1], buy1[i-1]+prices[i]) // 一次交易，卖出
	// buy2[i] = max(buy2[i-1], sell1[i-1]-prices[i]) // 二次交易，买入
	// sell2[i] = max(sell2[i-1], buy2[i-1]+prices[i]) // 二次交易，卖出
	l := len(prices)
	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0] // 买入卖出买入
	sell2 := 0         // 买入卖出买入卖出

	for i := 1; i < l; i++ {
		// 这里buy1、sell1、buy2没有存旧值，无论题目中是否允许「在同一天买入并且卖出」这一操作，最终的答案都不会受到影响，这是因为这一操作带来的收益为零
		buy1 = max(buy1, -prices[i])       // 一次交易，买入
		sell1 = max(sell1, buy1+prices[i]) // 一次交易，卖出
		buy2 = max(buy2, sell1-prices[i])  // 二次交易，买入
		sell2 = max(sell2, buy2+prices[i]) // 二次交易，卖出
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

