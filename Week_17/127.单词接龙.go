/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// bfs
	// 双向bfs
	dict := make(map[string]struct{})
	for _, word := range wordList {
		dict[word] = struct{}{}
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}

	q := []string{beginWord}
	q2 := []string{endWord}
	l := len(beginWord)
	visited := map[string]bool{beginWord: true}
	visited2 := map[string]bool{endWord: true}

	findLadder := func(word string) {
		arr := []byte(word)
		for i := 0; i < l; i++ {
			oldChar := arr[i]
			for j := byte('a'); j <= 'z'; j++ {
				if j == oldChar {
					continue
				}

				arr[i] = j
				curWord := string(arr)
				if _, ok := dict[curWord]; ok && !visited[curWord] {
					visited[curWord] = true
					q = append(q, curWord)
				}
			}
			arr[i] = oldChar
		}
	}

	step := 1
	for len(q) > 0 && len(q2) > 0 {
		if len(q) > len(q2) {
			q, q2 = q2, q
			visited, visited2 = visited2, visited
		}
		count := len(q)
		for i := 0; i < count; i++ {
			// contains
			for _, v := range q2 {
				if v == q[i] {
					return step
				}
			}
			findLadder(q[i])
		}
		step++
		q = q[count:]
	}
	return 0
}

// @lc code=end

