package _go

type ListNode struct {
	Val  int
	Next *ListNode
}

// 常数及空间复杂度，不能使用递归
func sortList(head *ListNode) *ListNode {
	// 获取整个链表的长度
	node := head
	l := 0
	for node != nil {
		l++
		node = node.Next
	}
	var l1, r1, l2, r2, start, end, lastTeamEnd, next *ListNode
	// 步长每次乘以二
	for i := 1; i < l; i <<= 1 {
		l1 = head
		r1 = findEnd(l1, i)
		// 因为有记录长度，这里l2不会为空，不用判断
		l2 = r1.Next
		r2 = findEnd(l2, i)
		// 记录下一组
		next = r2.Next
		r1.Next = nil
		r2.Next = nil
		head, lastTeamEnd = merage(l1, r1, l2, r2)
		for next != nil {
			l1 = next
			r1 = findEnd(l1, i)
			l2 = r1.Next
			if l2 == nil {
				lastTeamEnd.Next = l1
				break
			}
			r2 = findEnd(l2, i)
			next = r2.Next
			r1.Next = nil
			r2.Next = nil
			start, end = merage(l1, r1, l2, r2)
			lastTeamEnd.Next = start
			lastTeamEnd = end
		}

	}
	return head

}

// 寻找指定长度的链表区域
func findEnd(start *ListNode, k int) *ListNode {
	for start.Next != nil && k > 1 {
		start = start.Next
		k--
	}
	return start
}

// 归并(整合后返回头和尾)
func merage(l1, r1, l2, r2 *ListNode) (start, end *ListNode) {
	if l1.Val < l2.Val {
		start = l1
		l1 = l1.Next
	} else {
		start = l2
		l2 = l2.Next
	}
	node := start
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
			node = node.Next
		} else {
			node.Next = l2
			l2 = l2.Next
			node = node.Next
		}
	}

	if l1 == nil {
		node.Next = l2
		end = r2
	} else {
		node.Next = l1
		end = r1
	}
	return

}
