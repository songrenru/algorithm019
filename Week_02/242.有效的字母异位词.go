/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map1 := make(map[rune]int)
	for _, v := range s {
		map1[v]++
	}
	for _, v := range t {
		map1[v]--
	}
	for _, v := range map1 {
		if v != 0 {
			return false
		}
	}
	return true
}
// sort
// hash
// @lc code=end

