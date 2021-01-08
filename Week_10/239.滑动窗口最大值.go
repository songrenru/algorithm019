/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	// queque,进入一个大值，把前面的都顶掉
	push := func (v int)  {
		for len(q) > 0 && v > q[len(q) - 1] {
			q = q[:len(q) - 1]
		}
		q = append(q, v)
	}
	// 构造窗口
	for i := 0; i < k; i++ {
		push(nums[i])
	}

	// 开始滑动
	l := len(nums)
	res := []int{q[0]}
	for i := k; i < l; i++ {
		// 出值 = 最大值，最大值出栈
		if nums[i - k] == q[0] {
			q = q[1:]
		}
		// 进入一个大值，把前面的都顶掉
		push(nums[i])
		// 记录
		res = append(res, q[0])
	}
	return res
}
// @lc code=end

