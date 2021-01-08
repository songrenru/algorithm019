/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	l := len(nums)
	i := 0
	for j := 1; j < l; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return i + 1
}
// @lc code=end

