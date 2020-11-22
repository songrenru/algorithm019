/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return helper(x, n)
}

func helper(x float64, n int) float64 {
	// recursion terminator
	if n <= 1 {
		return x
	}
	// process current logic
	// dirll down
	re := helper(x, n/2)
	re *= re
	if n%2 == 1 {
		re *= x
	}
	return re
	// reverse state
}

// @lc code=end

