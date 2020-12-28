/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	strArr := strings.Split(s, " ")
	l := len(pattern)
	if l != len(strArr) {
		return false
	}
	
	// 双向连接，即恰好一一对应
	map1 := map[byte]string{}
	map2 := map[string]byte{}

	for i := 0; i < l; i++ {
		v1, ok1 := map1[pattern[i]]
		v2, ok2 := map2[strArr[i]]

		if !ok1 && !ok2 {
			map1[pattern[i]] = strArr[i]
			map2[strArr[i]] = pattern[i]
		} else if v1 != strArr[i] || v2 != pattern[i] {
			return false
		}
	}
	return true
}
// @lc code=end

