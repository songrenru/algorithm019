/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}

	helper := make([]int, k)
	copy(helper, nums[l - k:])

	for i := l - k - 1; i >= 0; i-- {
		nums[i + k] = nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = helper[i]
	}
}
// @lc code=end

