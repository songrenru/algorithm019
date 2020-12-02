/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	endReachable := len(nums) - 1
	for i := endReachable; i >= 0; i-- {
		if nums[i] + i >= endReachable  {
			endReachable = i
		}
	}
	return endReachable == 0
}
// @lc code=end

