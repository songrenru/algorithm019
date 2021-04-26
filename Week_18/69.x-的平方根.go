/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	// 牛顿迭代法
	// binary search
	if x <= 1 {
		return x
	}

	lo, hi := 1, x
	for lo <= hi {
		mid := (hi-lo)>>1 + lo
		tmp := mid * mid
		if tmp == x {
			return mid
		} else if tmp > x {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return hi
}

// @lc code=end

