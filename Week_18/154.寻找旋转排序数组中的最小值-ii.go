/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	// [2,2,2,0,0,0,1]
	// 1. 正常比较右边界
	// 2. 处理等于的特殊情况
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else if nums[mid] < nums[hi] {
			hi = mid
		} else { // nums[mid] == nums[hi]
			hi -= 1
		}
	}
	return nums[lo]
}

// @lc code=end

