/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := (hi-lo)>>1 + lo
		tmp := mid * mid
		if tmp == num {
			return true
		} else if tmp > num {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}

// @lc code=end

