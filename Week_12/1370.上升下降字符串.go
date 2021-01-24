/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

// @lc code=start
func sortString(s string) string {
	// 暴力，排序后扫描
	// 桶计数
	bucket := [26]int{}
	l := len(s)
	for i := 0; i < l; i++ {
		bucket[s[i] - 'a']++
	}

	res := []byte{}
	for len(res) < l {
		for i := 0; i < 26; i++ {
			if bucket[i] > 0 {
				res = append(res, byte(i + 'a'))
				bucket[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if bucket[i] > 0 {
				res = append(res, byte(i + 'a'))
				bucket[i]--
			}
		}
	}
	return string(res)
}
// @lc code=end

