/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	l1, l2 := len(s), len(t)
	if l1 != l2 {
		return false
	}
	bucket := [26]int{}
	for i := 0; i < l1; i++ {
		bucket[s[i]-'a']++
	}
	for i := 0; i < l2; i++ {
		bucket[t[i]-'a']--
	}
	for _, count := range bucket {
		if count != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

