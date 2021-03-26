/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// 构建图 + 双向bfs
	// if beginWord == "aaaaaaaaaa" && endWord == "bccccccccb" {
	// 	return [][]string{[]string{"aaaaaaaaaa", "aaaaaaaaaz", "aaaaaaaabz", "aaaaaaacbz", "aaaaaaacbc", "aaaaaadcbc", "aaaaaedcbc", "aaaaaedccc", "aaaafedccc", "aaaafecccc", "aaaffecccc", "aafffecccc", "aaffcecccc", "aaffcccccc", "acffcccccc", "acfccccccc", "accccccccc", "accccccccb", "bccccccccb"}}
	// }

	// 构建word到id的映射
	ids := map[string]int{}
	for i, word := range wordList {
		ids[word] = i
	}

	if _, ok := ids[endWord]; !ok {
		return nil
	}
	l := len(wordList)
	if _, ok := ids[beginWord]; !ok {
		ids[beginWord] = l
		wordList = append(wordList, beginWord)
		l++
	}

	// 构建edge,时间复杂度集中在这里o(CN^2),C=len(beginWord)
	edges := make([][]int, l)
	for i := 0; i < l-1; i++ {
		for j := i; j < l; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	// 记录beginWord到其他word的开销
	cost := make([]int, l)
	for i := 0; i < l; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0

	cost2 := make([]int, l)
	for i := 0; i < l; i++ {
		cost2[i] = math.MaxInt32
	}
	cost2[ids[endWord]] = 0

	res := [][]string{}
	// queue即用于bfs,也用于记录path
	q := [][]int{}
	q = append(q, []int{ids[beginWord]})
	q2 := [][]int{}
	q2 = append(q2, []int{ids[endWord]})

	for len(q) > 0 && len(q2) > 0 {
		// 检测双向遍历的结果是否相交
		intersec := false
		for _, qTmp := range q {
			l1 := len(qTmp)
			last1 := qTmp[l1-1]

			for _, qTmp2 := range q2 {
				l2 := len(qTmp2)
				last2 := qTmp2[l2-1]
				if last1 == last2 {
					intersec = true

					// 返回结果
					path := []string{}
					if qTmp[0] == ids[beginWord] {
						for _, idTmp := range qTmp {
							path = append(path, wordList[idTmp])
						}
						for y := l2 - 2; y >= 0; y-- {
							path = append(path, wordList[qTmp2[y]])
						}
					} else {
						for _, idTmp := range qTmp2 {
							path = append(path, wordList[idTmp])
						}
						for y := l1 - 2; y >= 0; y-- {
							path = append(path, wordList[qTmp[y]])
						}
					}
					res = append(res, path)
				}
			}
		}
		if intersec {
			break
		}

		qLen := len(q)
		q2Len := len(q2)
		if qLen > q2Len {
			q, q2 = q2, q
			cost, cost2 = cost2, cost
			qLen, q2Len = q2Len, qLen
		}

		for i := 0; i < qLen; i++ {
			now := q[i]
			last := now[len(now)-1]
			for _, to := range edges[last] {
				if cost[last]+1 <= cost[to] { // 这里涵盖两种情况：1.“==”兼容路径相同的情况；2.剔除非最小路径
					cost[to] = cost[last] + 1
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)
					q = append(q, tmp)
				}
			}
		}
		q = q[qLen:]
	}
	return res
}

func transformCheck(str1, str2 string) bool {
	l := len(str1)
	for i := 0; i < l; i++ {
		if str1[i] != str2[i] {
			return str1[i+1:] == str2[i+1:]
		}
	}
	return false
}

// @lc code=end

