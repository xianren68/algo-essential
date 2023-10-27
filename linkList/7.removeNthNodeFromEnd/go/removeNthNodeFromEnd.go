package _go
type ListNode struct {
	Val int
	Next *ListNode
	}
	func removeNthFromEnd(head *ListNode, n int) *ListNode {
		if head == nil {
			return nil
		}
		before,after := head,head
		// 前面指针先走n步
		i:=0
		for ;before != nil && i < n;i++{
			before = before.Next
		}
		if before == nil {
			if i < n {
				// 总长度不够n,返回原值
				return head
			}else {
				// 链表长度恰好为n,删除第一个节点
				return head.Next
			}
		}
		// 后面指针的前一个值
		var pre *ListNode
		// 两个指针同时走，前面的走到末尾，后面的正好到要删除的值上
		for before != nil {
			pre = after
			after = after.Next
			before = before.Next
		}
		// 删除指定位置节点
		pre.Next = after.Next
		return head
	}

