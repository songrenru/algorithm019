/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	// 暴力递归，最后检测括号有效性
	// 回溯 + 剪枝 // path直接用string传递，代码更简洁，但空间消耗更大
	res := []string{}
	var dfs func(left, right int, path []byte)
	dfs = func(left, right int, path []byte) {
		if left == n && right == n {
			temp := string(path)
			res = append(res, temp)
			return
		}
		if left < n {
			path = append(path, '(')
			dfs(left + 1, right, path)
			path = path[:len(path) - 1]
		}
	
		if right < left {
			path = append(path, ')')
			dfs(left, right + 1, path)
			// path = path[:len(path) - 1] // 右括号不存在回溯
		}
	}
	dfs(0, 0, []byte{})
	return res
}

// @lc code=end

