/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	lens := len(nums)
	if lens == 1 {
		return 0
	}
	count := 0
	endPoint := lens - 1
	nextEndPoint := lens - 2
	for nextEndPoint >= 0 {
		for i := nextEndPoint; i >= 0; i-- {
			// 找能一步跳到指定点的最远点
			if nums[i] + i >= endPoint {
				nextEndPoint = i // 最远点迁移
			}
		}
		endPoint = nextEndPoint
		nextEndPoint--
		count++
	}
	return count
}
// @lc code=end

