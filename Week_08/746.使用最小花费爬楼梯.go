/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	// dp
	// f(0) = 0, f(1) = min(cost[0], cost[1])
	// f(2) = min(f(2-1) + cost[2], f(2-2) + cost[2-1])
	// f(k) = min(f(k-1) + cost[k], f(k-2) + cost[k-1])
	lens := len(cost)
	dp := make([]int, lens)
	dp[0] = 0
	dp[1] = min(cost[0], cost[1])
	for i := 2; i < lens; i++ {
		// dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
		dp[i] = min(dp[i - 1] + cost[i], dp[i - 2] + cost[i - 1])
	}
	return dp[lens - 1]
}

func min(i, j int) int {
	if i < j { return i }
	return j
}
// @lc code=end

