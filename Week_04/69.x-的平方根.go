/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left, right := 1, x
	for left <= right {
		mid := left + (right - left) / 2 // 防止int求和溢出
		temp := mid * mid
		if temp == x {
			return mid
		} else if (temp > x) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
// @lc code=end

