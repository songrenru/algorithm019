/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	l := len(S)
	count := 0
	res := []byte{}
	for i := 0; i < l; i++ {
		if S[i] == ')' {
			count--
		}
		if count > 0 {
			res = append(res, S[i])
		}
		if S[i] == '(' {
			count++
		}
	}
	return string(res)
}

// @lc code=end

