/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
// func generateParenthesis(n int) []string {
// 	res := new([]string)
// 	realGenerator(n, n, "", res)
// 	return *res
// }

// func realGenerator(left, right int, s string, res *[]string) {
// 	// recursion
// 	if right == 0 && left == 0 {
// 		*res = append(*res, s)
// 		return
// 	}
// 	// process current logic
// 	// drill down
// 	if left > 0 {
// 		s1 := s + "("
// 		realGenerator(left - 1, right, s1, res)
// 	}
// 	if right > left {
// 		s2 := s + ")"
// 		realGenerator(left, right - 1, s2, res)
// 	}
// 	return 
// 	// reverse status
// }
func generateParenthesis(n int) []string {
	res := []string{}
	var g func(left, right int, s string)
	g = func (left, right int, s string) {
		if right == 0 && left == 0 {
			res = append(res, s)
			return
		}
		if left > 0 {
			g(left - 1, right, s + "(")
		}
		if right > left {
			g(left, right - 1, s + ")")
		}
		return 
	}
	g(n, n, "")
	return res
}
// @lc code=end

