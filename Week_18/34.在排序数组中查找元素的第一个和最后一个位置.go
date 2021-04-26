/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := binarySearch(nums, target)
	if nums[left] != target {
		return []int{-1, -1}
	}
	right := binarySearch(nums, target+1)
	if nums[right] != target {
		right -= 1
	}
	return []int{left, right}
}

func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1 // tips: hi == len(nums),可以更好的兼容
	for lo < hi {
		mid := (hi-lo)>>1 + lo
		if nums[mid] == target {
			hi = mid
		} else if target < nums[mid] {
			hi = mid - 1
		} else { // nums[mid] < target
			lo = mid + 1
		}
	}
	return lo
}

// @lc code=end

