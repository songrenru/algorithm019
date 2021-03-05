func getLeastNumbers(arr []int, k int) []int {
	l := len(arr)
	if l == 0 || k == 0 {
		return nil
	}
	pq := &maxHeap{}
	for i := 0; i < k; i++ {
		heap.Push(pq, arr[i])
	}

	heap.Init(pq)

	for i := k; i < l; i++ {
		if arr[i] < pq.Peek.(int) {
			heap.Pop()
			heap.Push(pq, arr[i])
		}
	}

	return *pq
}

type maxHeap []int

type maxHeap []int

func (a maxHeap) Len() int           { return len(a) }
func (a maxHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a maxHeap) Less(i, j int) bool { return a[i] > a[j] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 查看堆顶元素
func (h *maxHeap) Peek() interface{} {
	return (*h)[0]
}

