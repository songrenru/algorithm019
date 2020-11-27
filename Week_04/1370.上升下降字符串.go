/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

// @lc code=start
func sortString(s string) string {
	bucket := [26]int{}
	lens := len(s)
	for _, c := range s {
		bucket[byte(c) - 'a']++
	}

	ans := []byte{}
	for len(ans) < lens {
		for i := byte(0); i < 26; i++ {
			if bucket[i] > 0 {
				ans = append(ans, i + byte('a'))
				bucket[i]--
			}
		}
		for i := byte(26); i > 0; i-- {
			idx := i - 1
			if bucket[idx] > 0 {
				ans = append(ans, idx + byte('a'))
				bucket[idx]--
			}
		}
	}
	return string(ans)
}
// @lc code=end

