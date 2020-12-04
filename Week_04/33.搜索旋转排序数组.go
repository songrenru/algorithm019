/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	lo, hi := 0, len(nums) - 1
	for lo <= hi {
		mid := lo + (hi - lo) / 2
		if nums[mid] == target {
			return mid
		}

		// 前半段有序
		if nums[lo] <= nums[mid] { // 这里的=，排除了边界情况
			if target >= nums[lo] && target < nums[mid] { // 上面的nums[mid] == target，已经过滤了一种
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // 后半段有序（mid是旋转点，前后都有序）
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
// @lc code=end

