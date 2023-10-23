package _go

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	// 快慢指针
	slow := head
	quick := head
	for {
		if quick == nil || quick.Next == nil {
			return nil
		}
		slow = slow.Next
		quick = quick.Next.Next
		if quick == slow {
			break
		}
	}
	// 快指针返回头节点,变成慢指针
	quick = head
	for quick != slow {
		slow = slow.Next
		quick = quick.Next
	}
	return quick
}
