/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// 暴力递归
	// bfs
	// dp : 带记忆的递归（剪枝） ？ no,for
	// f(i) = min(f(i - k) in (coins)) + 1
	// dp(i) = min(dp[i - k] in (coins)) + 1
	dp := make([]int, amount + 1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ { // 存在无方案的情况
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i - coin] + 1) // 排除了无方案的case
			}
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
// @lc code=end

