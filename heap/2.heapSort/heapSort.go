package __heapSort

type heap struct {
	size int
	data []int
}

// 初始化heap
func initHeap(n int) *heap {
	h := &heap{
		size: -1,
		data: make([]int, n),
	}
	return h
}

// 插入数据
func (h *heap) push(v int) {
	h.size++
	i := h.size
	h.data[i] = v
	for i > 0 && h.data[i] < h.data[(i-1)/2] {
		// 与父节点交换位置
		h.data[i], h.data[(i-1)/2] = h.data[(i-1)/2], h.data[i]
		i = (i - 1) / 2
	}
}

// 弹出堆顶
func (h *heap) pop() int {
	ret := h.data[0]
	i := 0
	h.data[0] = h.data[h.size]
	h.size--
	l := i*2 + 1
	for l <= h.size {
		// 和左右子节点比较
		var best int
		if l+1 <= h.size && h.data[l] > h.data[l+1] {
			best = l + 1
		} else {
			best = l
		}
		if h.data[best] < h.data[i] {
			h.data[best], h.data[i] = h.data[i], h.data[best]
		}
		i = best
		l = i*2 + 1
	}
	return ret
}

func sortArray(nums []int) []int {
	h := initHeap(len(nums))
	for _, num := range nums {
		h.push(num)
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = h.pop()
	}
	return nums
}
