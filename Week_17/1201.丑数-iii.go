/*
 * @lc app=leetcode.cn id=1201 lang=golang
 *
 * [1201] 丑数 III
 */

// @lc code=start
func nthUglyNumber(n int, a int, b int, c int) int {
	// 预生成法
	// 最小堆预生成
	h := &IntHeap{a, b, c}
	heap.Init(h)
	res := make([]int, n+1)
	res[0] = 1
	count := 0
	for count < n {
		min := heap.Pop(h).(int)
		if min != res[count] {
			count++
			res[count] = min
			heap.Push(h, min*a)
			heap.Push(h, min*b)
			heap.Push(h, min*c)
		}
	}
	return res[n]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end

