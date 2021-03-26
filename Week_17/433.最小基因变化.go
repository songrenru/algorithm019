/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	// bfs
	// 双向bfs
	bankMap := make(map[string]struct{})
	for _, v := range bank {
		bankMap[v] = struct{}{}
	}
	if _, ok := bankMap[end]; !ok {
		return -1
	}

	maps := map[byte][3]byte{
		'A': [3]byte{'C', 'G', 'T'},
		'C': [3]byte{'A', 'G', 'T'},
		'G': [3]byte{'A', 'C', 'T'},
		'T': [3]byte{'A', 'C', 'G'},
	}

	step := 0
	l := len(start)
	q := []string{start}
	q2 := []string{end}
	for len(q) > 0 && len(q2) > 0 {
		if len(q) > len(q2) {
			q, q2 = q2, q
		}

		size := len(q)
		for i := 0; i < size; i++ {
			if contains(q[i], q2) {
				return step
			}

			// findWords
			cur := []byte(q[i])
			for j := 0; j < l; j++ {
				oldChar := cur[j]
				for _, c := range maps[oldChar] {
					cur[j] = c
					tmp := string(cur)

					if _, ok := bankMap[tmp]; ok { // 这里没有排除遍历过的，可能陷入无限循环
						q = append(q, tmp)
					}
				}
				cur[j] = oldChar
			}
		}

		step++
		q = q[size:]
	}
	return -1
}

func contains(str string, strS []string) bool {
	for _, v := range strS {
		if str == v {
			return true
		}
	}
	return false
}

// @lc code=end

