/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	res := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				res[j] = 1
			} else {
				res[j] = res[j] + res[j-1]
			}
		}
	}
	return res[n-1]
}
// @lc code=end

