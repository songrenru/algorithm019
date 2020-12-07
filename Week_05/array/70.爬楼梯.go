/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	// climbStairs(n) = climbStairs(n-1) + climbStairs(n-2)
	// 和斐波那契是一样的
	
	// 1. 暴力递归，时间log2^n,空间n
	// 2. 动态规划，求最优解
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
// @lc code=end

