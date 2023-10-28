package __mergeKSortedLists

type ListNode struct {
	Val  int
	Next *ListNode
}
type heap struct {
	size int
	data []*ListNode
}

// 初始化heap
func initHeap(n int) *heap {
	h := &heap{
		size: -1,
		data: make([]*ListNode, n),
	}
	return h
}

// 插入数据
func (h *heap) push(v *ListNode) {
	h.size++
	i := h.size
	h.data[i] = v
	for i > 0 && h.data[i].Val < h.data[(i-1)/2].Val {
		// 与父节点交换位置
		h.data[i], h.data[(i-1)/2] = h.data[(i-1)/2], h.data[i]
		i = (i - 1) / 2
	}
}

// 弹出堆顶
func (h *heap) pop() *ListNode {
	ret := h.data[0]
	i := 0
	h.data[0] = h.data[h.size]
	h.size--
	l := i*2 + 1
	for l <= h.size {
		// 和左右子节点比较
		var best int
		if l+1 <= h.size && h.data[l].Val > h.data[l+1].Val {
			best = l + 1
		} else {
			best = l
		}
		if h.data[best].Val < h.data[i].Val {
			h.data[best], h.data[i] = h.data[i], h.data[best]
		}
		i = best
		l = i*2 + 1
	}
	return ret
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 初始化小根堆
	h := initHeap(len(lists))
	// 将每个链表头节点入堆
	for _, list := range lists {
		if list != nil {
			h.push(list)
		}
	}
	// 全是空链表
	if h.size == -1 {
		return nil
	}
	// 头节点
	head := h.pop()
	if head.Next != nil {
		h.push(head.Next)
	}
	node := head
	for h.size >= 0 {
		n := h.pop()
		node.Next = n
		node = node.Next
		if n.Next != nil {
			h.push(n.Next)
		}
	}
	
	return head
}
