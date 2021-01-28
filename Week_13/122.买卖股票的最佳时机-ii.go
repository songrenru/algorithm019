/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// greddy
	// dp
	// dp + 空间优化
	l := len(prices)
	if l <= 1 {
		return 0
	}
	
	cash := 0
	hold := -prices[0]
	for i := 1; i < l; i++ {
		cash, hold = max(cash, hold + prices[i]), max(hold, cash - prices[i])
	}
	return cash
}

func max(i, j int) int {
	if i  > j { return i }
	return j
}
// @lc code=end

