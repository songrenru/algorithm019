/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	j := 0
	lens := len(nums)
	for i := 0; i < lens; i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
// @lc code=end

