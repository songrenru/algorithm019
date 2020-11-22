/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	pre := []int{}
	for i := 0; i <= rowIndex; i++ {
		cur  := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				cur = append(cur, 1)
			} else {
				cur = append(cur, pre[j-1] + pre[j])
			}
		}
		pre = cur
	}
	return pre
}
// @lc code=end

