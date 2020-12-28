/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	// 构造hash table
	bankMap := map[string]int{}
	for i, v := range bank {
		bankMap[v] = i
	}

	// check
	if _, ok := bankMap[end]; !ok {
		return -1
	}

	// bfs
	mutation := [4]string{"A", "C", "G", "T"}
	queue := []string{start}
	used := make([]bool, len(bank))
	step := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			curr := queue[i]
			if curr == end {
				return step
			}
			// 逐个替换每个位的字母，比对bank
			for j := 0; j < 8; j++ {
				for k := 0; k < 4; k++ {
					temp := curr[:j] + mutation[k] + curr[j+1:]
					if idx, ok := bankMap[temp]; ok && !used[idx] {
						queue = append(queue, temp)
						used[idx] = true
					}
				}
			}
		}
		queue = queue[l:]
		step++
	}
	return -1
}
// @lc code=end

