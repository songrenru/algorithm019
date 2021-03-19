/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	// 快速幂 + 递归
	// 快速幂 + 迭代
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	res := 1.0
	x_contribute := x
	for n > 0 {
		if n&1 == 1 {
			res *= x_contribute
		}
		x_contribute *= x_contribute
		n >>= 1
	}
	return res
}

// @lc code=end

