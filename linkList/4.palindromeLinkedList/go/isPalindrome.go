package _go

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 快慢指针走到中点
	// 记录长度
	slow, quick := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
	}
	var reHead *ListNode
	reHead = reverse(slow.Next)
	rNode, node := reHead, head
	for rNode != nil && node != nil && rNode.Val == node.Val {
		rNode = rNode.Next
		node = node.Next
	}
	return node == nil || rNode == nil
}

// 反转链表
func reverse(start *ListNode) *ListNode {
	var cur, pre, next *ListNode
	cur = start
	for cur != nil {
		next = cur.Next
		// 将指针指向上一个节点
		cur.Next = pre
		pre = cur
		// 到下一个节点
		cur = next
	}
	return pre
}
