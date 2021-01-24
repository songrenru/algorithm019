/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 二次遍历 + dp
	// 一次遍历 + dp
	l := len(prices)

	minPrice := prices[0]
	profit := 0
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
	if i > j { return i }
	return j
}
// @lc code=end

