/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	// 预生成法
	// 最小堆预生成
	// dp预生成,即三指针
	dp := []int{1}
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < 1690; i++ {
		num1, num2, num3 := 2*dp[p2], 3*dp[p3], 5*dp[p5]
		uglyNum := min(num1, num2, num3)
		dp = append(dp, uglyNum)
		if num1 == uglyNum {
			p2++
		}
		if num2 == uglyNum {
			p3++
		}
		if num3 == uglyNum {
			p5++
		}
	}
	return dp[n-1]
}

func min(i, j, k int) int {
	if i < j {
		j = i
	}
	if j < k {
		return j
	}
	return k
}

// @lc code=end

