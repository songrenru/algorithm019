/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	l_s, l_g := len(s), len(g)
	i, j := 0, 0
	for i < l_s && j < l_g {
		if s[i] >= g[j] {
			j++
		}
		i++
	}
	return j
}

// @lc code=end

