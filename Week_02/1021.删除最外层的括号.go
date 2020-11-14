/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	res := make([]byte, 0)
	count := 0
    lens := len(S)
    for i := 0; i < lens; i++ {
		if S[i] == '(' {
			if count > 0 {
				res = append(res, S[i])
			}
			count++
		} else if S[i] == ')' {
			count--
			if count > 0 {
				res = append(res, S[i])
			}
		}
	}

	return string(res)
}
// @lc code=end

