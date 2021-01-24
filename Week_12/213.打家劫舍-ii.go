/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	// 1偷，n不偷
	// 1不偷，n偷
	return max(helper(nums[:l - 1]), helper(nums[1:]))
}

func helper(nums []int) int {
	// dp
	// dp[i][0] = max(dp[i - 1][0], dp[i - 1][1])
	// dp[i][1] = dp[[i - 1][0] + nums[i]
	l := len(nums)
	if l == 0 {
		return 0
	}
	dp := make([][2]int, l)
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1])
		dp[i][1] = dp[i - 1][0] + nums[i]
	}
	return max(dp[l - 1][0], dp[l - 1][1])
}

func max(i, j int) int {
	if i > j { return i }
	return j
}
// @lc code=end

