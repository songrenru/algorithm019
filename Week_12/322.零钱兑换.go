/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// dp: f(i) = min(f(i - coin)) + 1 // coin in coins
	// bfs
	// 完全背包
	dp := make([]int, amount + 1)
	// dp[0] = 0
	for i := 1; i <= amount; i++ { dp[i] = amount + 1 }

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i - coin] + 1, dp[i])
			}
		}
	}
	if dp[amount] > amount { return -1 }
	return dp[amount]
}

func min(i, j int) int {
	if i < j { return i }
	return j
}
// @lc code=end

