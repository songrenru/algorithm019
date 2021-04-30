/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// dp[i] = min(dp[i-k]) + 1, k in coins
	// bfs
	if amount == 0 {
		return 0
	}

	visited := map[int]bool{amount: true}
	q := []int{amount}
	step := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[i]

			for _, coin := range coins {
				next := cur - coin
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
		q = q[size:]
	}
	return -1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

