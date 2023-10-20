package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var heap []int
var size int

// 初始化堆
func initHeap(n int) {
	heap = make([]int, n)
}

// 添加
func push(n int) {
	heap[size] = n
	size++
	// 向上替换
	i := size - 1
	for i > 0 && heap[(i-1)/2] > heap[i] {
		heap[(i-1)/2], heap[i] = heap[i], heap[(i-1)/2]
		i = (i - 1) / 2
	}
}

// 弹出
func pop() int {
	ret := heap[0]
	heap[0] = heap[size-1]
	size--
	i := 0
	l := i*2 + 1
	var best int
	for l < size {
		if l+1 < size && heap[l] > heap[l+1] {
			best = l + 1
		} else {
			best = l
		}
		if heap[i] > heap[best] {
			heap[i], heap[best] = heap[best], heap[i]
			i = best
			l = i*2 + 1
		} else {
			break
		}
	}
	return ret
}

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func read() int {
	// 指示是否有-号
	flag := false
	ret := 0
	for c, _ := in.ReadByte(); (c >= '0' && c <= '9') || c == '-'; c, _ = in.ReadByte() {
		if c == '-' {
			flag = true
			continue
		}
		ret = (ret * 10) + int(c-'0')
	}
	if flag {
		ret = -ret
	}
	return ret
}
func computed(nums [][2]int) int {
	// 将线段按照起始位置排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})
	ans := 0
	for _, num := range nums {
		// 前面线段的结束位置 > 当前线段的起始位置，有重叠
		// 将不重叠的出堆
		for size > 0 && heap[0] <= num[0] {
			pop()
		}
		// 加入当前线段的结尾
		push(num[1])
		if size > ans {
			ans = size
		}
	}
	return ans
}

func main() {
	n := read()
	nums := make([][2]int, n)
	initHeap(n)
	for i := 0; i < n; i++ {
		nums[i] = [2]int{read(), read()}
	}
	out.WriteString(strconv.Itoa(computed(nums)))
	out.Flush()
}
