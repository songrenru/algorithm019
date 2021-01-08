/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	l := len(nums)
	k %= l
	reverse(nums, 0, l - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, l - 1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
// @lc code=end

