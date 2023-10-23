package __median

type MedianFinder struct {
	big   *heap // 大根堆 保存较小的一部分
	small *heap // 小根堆 保存较大的一部分
}

// 定义比较器方法
type compareFunc func(i, j int) bool

// 堆
type heap struct {
	data [100000]int
	size int
	fn   compareFunc
}

// 初始化堆
func initHeap(fn compareFunc) *heap {
	return &heap{
		fn: fn,
	}
}

// 入堆
func (h *heap) push(num int) {
	h.data[h.size] = num
	i := h.size
	h.size++
	for i > 0 && h.fn(h.data[(i-1)/2], h.data[i]) {
		h.data[(i-1)/2], h.data[i] = h.data[i], h.data[(i-1)/2]
		i = (i - 1) / 2
	}
}

// 出堆
func (h *heap) pop() int {
	h.size--
	ret := h.data[0]
	h.data[0] = h.data[h.size]
	i := 0
	l := i*2 + 1
	for l < h.size {
		best := 0
		// 比较左右子节点
		if (l+1 < h.size) && h.fn(h.data[l], h.data[l+1]) {
			best = l + 1
		} else {
			best = l
		}
		if h.fn(h.data[i], h.data[best]) {
			h.data[i], h.data[best] = h.data[best], h.data[i]
			i = best
			l = i*2 + 1
		} else {
			break
		}

	}
	return ret
}

func Constructor() MedianFinder {
	// 建立两个堆，每次取中位数时取出它们的堆顶即可
	return MedianFinder{
		big: initHeap(func(i, j int) bool {
			return i < j
		}),
		small: initHeap(func(i, j int) bool {
			return i > j
		}),
	}

}

func (m *MedianFinder) AddNum(num int) {
	//// 判断这个数是不是比大根堆的堆顶还大
	//if m.big.size == 0 || num > m.big.data[0] {
	//	// 加入小根堆
	//	m.small.push(num)
	//} else {
	//	m.big.push(num)
	//}
	//// 保持平衡
	//// 判断两个堆的长度相差是否大于1
	//if m.small.size-m.big.size > 1 {
	//	m.big.push(m.small.pop())
	//} else if m.big.size-m.small.size > 1 {
	//	m.small.push(m.big.pop())
	//}
	if m.big.size == m.small.size {
		// 加入大根堆
		m.big.push(num)
		// 将大根堆顶给小根堆
		m.small.push(m.big.pop())
	} else {
		m.small.push(num)
		m.big.push(m.small.pop())
	}
}

func (m *MedianFinder) FindMedian() float64 {
	//// 判断两个堆的大小
	//if m.big.size > m.small.size {
	//	return float64(m.big.data[0])
	//} else if m.big.size < m.small.size {
	//	return float64(m.small.data[0])
	//} else {
	//	// 两边相等
	//	return float64(m.small.data[0]+m.big.data[0]) / 2
	//}
	if m.big.size == m.small.size {
		return float64(m.small.data[0]+m.big.data[0]) / 2
	}
	return float64(m.small.data[0])
}
