/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	// 双指针
	lens := len(nums)
	positionNotZero := 0
	for i := 0; i < lens; i++ {
		if nums[i] != 0 { // && positionNotZero != i，[2, 1]失败
			if i != positionNotZero {
				nums[positionNotZero], nums[i] = nums[i], nums[positionNotZero]
			}
			positionNotZero++
		}
	}
}
// @lc code=end

