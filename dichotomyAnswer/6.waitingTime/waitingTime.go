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

// 堆方法
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
	// 到自己时取出第一个服务员所记录的时间
	return heap.Pop(&h).([2]int)[0]
}

// 二分答案法
func waitTime2(nums []int, n int) int {
	// 等待的时间范围 0 ~ 耗时最少服务员的时间*等待的人数
	// 在这个范围不断二分，寻找能服务k个客人的时间
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	// 二分
	ans := 0
	l, r := 0, n*min
	var m int
	for l <= r {
		m = l + (r-l)>>1
		// 当前时间要让自己也得到服务，所以n+1
		if f(nums, m) >= n+1 {
			// 缩小等待时间，记录答案
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans

}

// 返回在当前时间内所能接待的客人数
func f(nums []int, m int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		// 一个服务员可以接待的客人
		res += (m / nums[i]) + 1
	}
	return res
}

func main() {
	fmt.Println(waitTime([]int{2, 4, 6, 4, 3, 1}, 10))
	fmt.Println(waitTime2([]int{2, 4, 6, 4, 3, 1}, 10))
}
