/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	// 大顶堆，最坏情况O(NlogN)
	// 单调队列,O(n)
	if k == 1 {
		return nums
	}

	// 构建单调队列
	q := []int{math.MinInt32}
	for i := 0; i < k; i++ {
		// 打掉前面比自己小的元素
		for j := len(q) - 1; j >= 0 && nums[i] > q[j]; j-- {
			q = q[0:j]
		}
		q = append(q, nums[i])
	}

	// 开始窗口滑动
	res := []int{q[0]}
	l := len(nums)
	for i := k; i < l; i++ {
		// 判定划出窗口的首元素，是否是最大值
		outIdx := i - k
		if nums[outIdx] == q[0] {
			q = q[1:]
		}

		// 打掉前面比自己小的元素
		for j := len(q) - 1; j >= 0 && nums[i] > q[j]; j-- {
			q = q[0:j]
		}
		q = append(q, nums[i])
		res = append(res, q[0])
	}

	return res
}

// @lc code=end

