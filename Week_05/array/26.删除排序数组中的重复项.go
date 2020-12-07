/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	cursor := 0
	lens = len(nums)
	for i := 1; i < lens; i++ {
		if nums[i] != nums[cursor] {
			cursor++
			nums[cursor] = nums[i]
		}
	}
	return cursor + 1
}
// @lc code=end

