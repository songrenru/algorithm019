/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// bfs
	// 双向bfs
	dict := map[string]bool{}
	for _, word := range wordList {
		dict[word] = true
	}
	if !dict[endWord] {
		return 0
	}
	visited := map[string]bool{}
	queueLeft := []string{beginWord}
	queueRight := []string{endWord}
	l := len(beginWord)

	findLadder := func(word string) bool {
		arr := []byte(word)
		for i := 0; i < l; i++ {
			oldChar := arr[i]
			for j := byte('a'); j <= 'z'; j++ {
				if j == oldChar {
					continue
				}
				curChar := j
				arr[i] = curChar
				curWord := string(arr)

				if dict[curWord] {
					for _, q := range queueRight {
						if q == curWord {
							return true
						}
					}
					if !visited[curWord] {
						visited[curWord] = true
						queueLeft = append(queueLeft, curWord)
					}
				}

				arr[i] = oldChar
			}
		}
		return false
	}

	step := 1
	for len(queueLeft) > 0 && len(queueRight) > 0 {
		if len(queueLeft) > len(queueRight) {
			queueLeft, queueRight = queueRight, queueLeft
		}
		count := len(queueLeft)
		for i := 0; i < count; i++ {
			if findLadder(queueLeft[i]) {
				return step + 1
			}
		}
		step++
		queueLeft = queueLeft[count:]
	}
	return 0
}
// @lc code=end

