/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
var mem map[int]int = make(map[int]int)
func fib(N int) int {
	// 暴力递归
	// 状态记忆的递归 // 剪枝
	// dp 递推方程：f(N) = f(N-1) + f(N-2)
	/*
	 * 1. top to down，状态记忆递归
	 * 2. down to top, for形式，滚动数组，空间最优
	 */
	if N <= 1 {
		return N
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= N; i++ {
		p, q = q, r
		r = p + q
	}
	return r
}
// @lc code=end

