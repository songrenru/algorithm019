/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
type pair struct {
	ch byte
	pos int
}

func firstUniqChar(s string) int {
	// hashTable + 队列,延迟删除
	l := len(s)
	pos := [26]int{}
	for i := range pos {
		pos[i] = l
	}

	queue := []pair{}
	for i := 0; i < l; i++ {
		idx := s[i] - 'a'
		if pos[idx] == l {
			pos[idx] = i
			queue = append(queue, pair{idx, i})
		} else {
			pos[idx] = l + 1
			// if len(queue) > 0 && pos[queue[0].ch] == l + 1 {
			// 	queue = queue[1:]
			// }
			for len(queue) > 0 && pos[queue[0].ch] == l + 1 { // 把所有重复项打掉
				queue = queue[1:]
			}
		}	
	}

	if len(queue) > 0 {
		return queue[0].pos
	}
	return -1


	// hashTable思想，记录索引
	// l := len(s)
	// pos := [26]int{}
	// for i := 0; i < l; i++ {
	// 	idx := s[i] - 'a'
	// 	if pos[idx] == 0 {
	// 		pos[idx] = i + 1 // 索引从1开始，和0区分开
	// 	} else if pos[idx] >= 1 {
	// 		pos[idx] = -1
	// 	}
	// }

	// first := l + 1
	// for i := 0; i < 26; i++ {
	// 	if pos[i] >= 1 && pos[i] < first {
	// 		first = pos[i]
	// 	}
	// }

	// if first == l + 1 {
	// 	return -1
	// }
	// return first - 1

	// hashTable思想，char计数
	// l := len(s)
	// counter := [26]int{}
	// for i := 0; i < l; i++ {
	// 	counter[s[i]  -  'a']++
	// }
	// for i := 0; i < l; i++ {
	// 	if counter[s[i] - 'a'] == 1 {
	// 		return i
	// 	}
	// }
	// return -1
}
// @lc code=end

