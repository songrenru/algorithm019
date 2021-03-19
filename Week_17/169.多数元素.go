/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	// hash
	// sort
	// 随机化，理论时间复杂度无穷，众数随机期望次数2，时间复杂度O(n)
	// 分治，需要证明
	// Boyer-Moore 投票算法(最优)
	l := len(nums)
	count := 0
	candidate := 0
	for i := 0; i < l; i++ {
		if count == 0 {
			candidate = nums[i]
			count++
		} else if count == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// @lc code=end

