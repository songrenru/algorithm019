/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	q := &hp{}
	heap.Init(q)
	// 构造窗口
	for i := 0; i < k; i++ {
		heap.Push(q, nums[i])
	}

	// 开始滑动
	l := len(nums)
	res := []int{(*q)[0]}
	for i := k; i < l; i++ {
		// 出值 = 最大值，最大值出栈
		// if nums[i - k] == (*q)[0] {
		// 	heap.Pop(q)
		// }
		heap.Push(q, nums[i])
		for (*q)[0] <= nums[i - k] {
            heap.Pop(q)
        }
		// 进入一个大值，把前面的都顶掉
		// heap.Push(q, nums[i])
		// 记录
		res = append(res, (*q)[0])
	}
	return res
}

type hp []int

func (a hp) Len() int           { return len(a) }
func (a hp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hp) Less(i, j int) bool { return a[i] > a[j] }
func (a *hp) Push(x interface{}) { *a = append(*a, x.(int)) }
func (a *hp) Pop() interface{} { 
	old := *a
	l := len(old)
	v := old[l - 1]
	*a = old[:l - 1]
	return v
}

// @lc code=end

