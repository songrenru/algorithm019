/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (hi-lo)>>1 + lo
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			if nums[lo] <= nums[mid] && nums[lo] > target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // nums[mid] < target
			if nums[mid] <= nums[hi] && nums[hi] < target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}
	return -1
}

// @lc code=end

