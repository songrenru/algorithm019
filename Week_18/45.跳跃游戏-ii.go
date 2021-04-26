/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	lastIdx := len(nums) - 1
	count := 0
	for lastIdx > 0 {
		last := lastIdx
		for i := lastIdx - 1; i >= 0; i-- {
			if nums[i]+i >= lastIdx {
				last = i
			}
		}
		lastIdx = last
		count++
	}

	return count
}

// @lc code=end

