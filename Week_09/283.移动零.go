/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	// 双指针
	cursor := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			nums[cursor], nums[i] = nums[i], nums[cursor]
			cursor++
		}
	}
}
// @lc code=end

