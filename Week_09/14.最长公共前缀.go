/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// 横向扫描(本质还是纵向扫描)
	// 纵向扫描
	l := len(strs)
	if l == 0 {
		return ""
	}

	lcs := strs[0]
	for i := 1; i < l; i++ {
		lcs = getLCS(lcs, strs[i])
		if lcs == "" {
			return ""
		}
	}
	return lcs
}

func getLCS(str1, str2 string) string {
	minLen := min(len(str1), len(str2))
	idx := 0
	for idx < minLen && str1[idx] == str2[idx] {
		idx++
	}
	return string(str1[:idx])
}

func min(i, j int) int {
	if i < j { return i }
	return j
}
// @lc code=end

