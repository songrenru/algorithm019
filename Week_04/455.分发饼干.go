/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	lenG := len(g)
	lenS := len(s)

	j := 0
	for i := 0; i < lenS && lenG != 0; i++ {
		if s[i] >= g[j] {
			j++
			lenG--
		}
	}
	return j
}
// @lc code=end

