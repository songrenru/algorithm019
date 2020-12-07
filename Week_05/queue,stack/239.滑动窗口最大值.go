/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	// 双端队列
	if nums == nil || k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}

	queue := []int{nums[0]}
	// 构建初始窗口
	for i := 1; i < k; i++ {
		for j := len(queue) - 1; j >= 0 && nums[i] > queue[j]; j-- {
			queue = queue[:j]
		}
		queue = append(queue, nums[i])
	}

	res := []int{queue[0]}
	lastIdx := len(nums) - 1
	for i := k; i <= lastIdx; i++ {
		oldMax := queue[0]
		for j := len(queue) - 1; j >= 0 && nums[i] > queue[j]; j-- {
			queue = queue[:j]
		}
		queue = append(queue, nums[i])
		if oldMax == nums[i-k] && oldMax == queue[0]{
			queue = queue[1:]
		}
		res = append(res, queue[0])
	}
	return res
}
// @lc code=end

