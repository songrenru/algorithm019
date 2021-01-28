/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp + 双循环
	// dp + 一遍循环
	l := len(prices)
	if l <= 1 {
		return 0
	}
	profit := 0
	minPrice := prices[0]
	for i := 1; i < l; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] > minPrice {
			profit = max(profit, prices[i] - minPrice)
		}
	}
	return profit
}

func max(i, j int) int {
	if i  > j { return i }
	return j
}
// @lc code=end

