package _go

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	end := teamEnd(start, k)
	if end == nil {
		return start
	}
	// 记录头节点
	head = end
	// 反转链表
	reverse(start, end)
	// start为上一组的尾节点
	lastTeamEnd := start
	for lastTeamEnd.Next != nil {
		start = lastTeamEnd.Next
		end = teamEnd(start, k)
		if end == nil {
			return head
		}
		reverse(start, end)
		lastTeamEnd.Next = end
		lastTeamEnd = start
	}
	return head
}

// 寻找到一组中的末尾节点
func teamEnd(start *ListNode, k int) (end *ListNode) {
	for start != nil && k != 1 {
		start = start.Next
		k--
	}
	end = start
	return
}

// 反转链表
func reverse(start, end *ListNode) {
	end = end.Next
	var pre, next, cur *ListNode
	cur = start
	for cur != end {
		next = cur.Next
		cur.Next = pre
		// 记录上一个节点
		pre = cur
		cur = next
	}
	// 连接下一组的节点
	start.Next = end
}
