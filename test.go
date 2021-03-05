package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func main() {
	h := &IntHeap{2, 3, 5}
	heap.Init(h)
	res := [1690]int{1}
	n := 0
	for n < 1689 {
		min := heap.Pop(h).(int)
		if min != res[n] {
			n++
			res[n] = min
			heap.Push(h, min*2)
			heap.Push(h, min*3)
			heap.Push(h, min*5)
		}
	}

	fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(res), "[]"), " ", ",", -1))
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
