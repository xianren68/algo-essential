package main

import (
	"container/heap"
	"fmt"
)

type Heap [][2]int

func (h *Heap) Less(i, j int) bool {
	return (*h)[i][0] < (*h)[j][0]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x any) {
	val := x.([2]int)
	*h = append(*h, val)
}

func (h *Heap) Pop() any {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

func (h *Heap) Len() int {
	return len(*h)
}
func waitTime(nums []int, n int) int {
	// 建立小根堆，每一项存储（服务员完成工作的时间，完成工作所需的时间）
	var h Heap = make(Heap, len(nums))
	for i, val := range nums {
		h[i] = [2]int{0, val}
	}
	heap.Init(&h)
	// 遍历所有用户，每次取出堆顶的元素，让他服务指定时间并再次入堆
	for i := 0; i < n; i++ {
		val := heap.Pop(&h).([2]int)
		val[0] += val[1]

		heap.Push(&h, val)
	}
	fmt.Println(h)
	// 到自己时取出第一个服务员所记录的时间
	return heap.Pop(&h).([2]int)[0]
}
func main() {
	fmt.Println(waitTime([]int{2, 4, 6, 4, 3, 1}, 10))
}
