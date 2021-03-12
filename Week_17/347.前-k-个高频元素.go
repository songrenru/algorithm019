/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	// min heap

	// 桶计数
	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}

	// 构建最小堆
	h := &EleHeap{}
	fre := 0
	for num, count := range counter {
		if fre < k {
			*h = append(*h, ele{num, count})
		} else {
			if fre == k {
				heap.Init(h)
			}
			if count > h.Peak().count {
				heap.Pop(h)
				heap.Push(h, ele{num, count})
			}
		}
		fre++
	}

	// 返回结果
	res := []int{}
	for _, e := range *h {
		res = append(res, e.num)
	}
	return res
}

type ele struct {
	num, count int
}

type EleHeap []ele

func (h EleHeap) Len() int           { return len(h) }
func (h EleHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h EleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h EleHeap) Peak() ele          { return h[0] }
func (h *EleHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(ele))
}
func (h *EleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
