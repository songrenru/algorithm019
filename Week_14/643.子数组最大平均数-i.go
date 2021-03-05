/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	l := len(nums)
	if l < k || k == 0 {
		return 0
	}

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	max := sum
	for i := k; i < l; i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

// @lc code=end
