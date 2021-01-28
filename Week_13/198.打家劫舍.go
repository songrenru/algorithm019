/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	// 二维dp
	// 一维dp
	// dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[0], nums[1])
	}
	pre, cur := nums[0], max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		pre, cur = cur, max(cur, pre + nums[i])
	}
	return cur
}

func max(i, j int) int {
	if i  > j { return i }
	return j
}
// @lc code=end

