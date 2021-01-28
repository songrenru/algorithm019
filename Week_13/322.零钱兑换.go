/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// dp[i] = min(dp[i - coin] + 1, dp[i]) // coin in coins && i >= coin
	// bfs
	if amount == 0 {
		return 0
	}
	l := len(coins)
	q := []int{amount}
	visited := map[int]bool{amount:true}
	step := 0
	for len(q) > 0 {
		count := len(q)
		for i := 0; i < count; i++ {
			curAmount := q[i]
			for j := l - 1; j >= 0; j-- {
				next := curAmount - coins[j]
				if next == 0 {
					return step + 1
				}
				if next > 0 && !visited[next] {
					visited[next] = true
					q = append(q, next)
				}
			}
		}
		step++
		q = q[count:]
	}
	return -1
}

func min(i, j int) int {
	if i < j { return i }
	return j
}
// @lc code=end

