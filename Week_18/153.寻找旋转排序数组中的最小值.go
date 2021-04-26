/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	// 比较left、mid
	// 比较right、mid
	lo, hi := 0, len(nums)-1
	for lo < hi { // lo < hi，就已经排除了mid == hi的情况
		mid := (hi-lo)>>1 + lo
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else { // nums[mid] <= nums[right]  // lo < hi，就已经排除了mid == hi的情况，所以这里只有nums[mid] < nums[right]
			hi = mid
		}
	}
	return nums[lo]
}

// @lc code=end

