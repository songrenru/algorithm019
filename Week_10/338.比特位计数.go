/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	res := make([]int, num + 1)
	for i := 1; i <= num; i++ {
		if i & 1 == 1 {
			res[i] = res[i - 1] + 1
		} else {
			res[i] = res[i / 2]
		}
	}
	return res
}
// @lc code=end

