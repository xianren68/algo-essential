package _go

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lA, endA := findEnd(headA)
	lB, endB := findEnd(headB)
	if endB == nil || endA == nil {
		return nil
	}
	if lA > lB {
		return findFirst(headA, headB, lA, lB)
	}
	return findFirst(headB, headA, lB, lA)
}

// 返回链表的长度及其尾节点
func findEnd(h *ListNode) (length int, end *ListNode) {
	if h == nil {
		return
	}
	for h.Next != nil {
		length++
		h = h.Next
	}
	length++
	end = h
	return
}

// 遍历指定位置
func findFirst(h1, h2 *ListNode, l1, l2 int) *ListNode {
	// 先走
	for l1 > l2 {
		l1--
		h1 = h1.Next
	}
	// 同时走，相同时为相交节点
	for h1 != h2 {
		h1 = h1.Next
		h2 = h2.Next
	}
	return h1
}
