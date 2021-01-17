/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	l1 := len(s)
	l2 := len(t)
	if l1 != l2 {
		return false
	}

	map1 := [26]int{}
	for i := 0; i < l1; i++ {
		map1[s[i] - 'a']++
	}

	for i := 0; i < l2; i++ {
		map1[t[i] - 'a']--
	}

	for _, v := range map1 {
		if v != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

