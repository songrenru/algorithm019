/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 字典
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	delete(wordSet, beginWord)

	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true

	var changeWordEveryOneLetter = func(currentWord string) bool {
		wordArr := []byte(currentWord)
		size := len(wordArr)

		for i := 0; i < size; i++ {
			originChar := wordArr[i]
			for c := byte('a'); c <= 'z'; c++ {
				if c == originChar {
					continue
				}

				wordArr[i] = c
				newWord := string(wordArr)
				if _, ok := wordSet[newWord]; ok {
					if newWord == endWord {
						return true
					}
					if !visited[newWord] {
						queue = append(queue, newWord)
						visited[newWord] = true
					}
				}
			}
			wordArr[i] = originChar
		}
		return false
	}

	// bfs
	step := 1 // 包含起点
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			currentWord := queue[i]
			if changeWordEveryOneLetter(currentWord) {
				return step + 1
			}
		}
		queue = queue[size:]
		step++
	}
	return 0
}
// @lc code=end

