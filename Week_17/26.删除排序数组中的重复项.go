/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	cursor := 0
	for i := 1; i < l; i++ {
		if nums[i] != nums[cursor] {
			cursor++
			nums[cursor] = nums[i]
		}
	}
	return cursor + 1
}

// @lc code=end

