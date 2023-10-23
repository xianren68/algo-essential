package _go

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 给每个节点产生一份复制节点
	node := head
	for node != nil {
		cNode := &Node{
			Val:  node.Val,
			Next: node.Next,
		}
		node.Next = cNode
		node = cNode.Next
	}
	// 将随机指针复制
	node = head
	for node != nil {
		// 判断当前节点有无random
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
		node = node.Next.Next
	}
	// 拆散两个链表
	cHead := head.Next
	node = head
	var next *Node
	for node != nil {
		// 记录下一个节点
		next = node.Next.Next
		// 复制节点指向复制节点
		if next == nil {
			node.Next.Next = nil

		} else {
			node.Next.Next = next.Next
		}
		node.Next = next
		node = next
	}
	return cHead
}
