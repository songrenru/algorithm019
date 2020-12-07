/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]--
		if digits[i] >= 0 {
			return digits
		}
		digits[i] = 9
	}
	return digits[1:]
}
// @lc code=end

