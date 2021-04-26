/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	// [2,3,1,1,4]
	endReachable := len(nums) - 1
	for i := endReachable - 1; i >= 0; i-- {
		if i+nums[i] >= endReachable {
			endReachable = i
		}
	}
	return endReachable == 0
}

// @lc code=end

