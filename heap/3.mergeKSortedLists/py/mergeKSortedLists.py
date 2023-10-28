from typing import Optional,List
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    # 堆数据区
    data:List[ListNode] = []
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        # 全是空链表
        if not lists:
            return None
        # 链表全入堆"
        for i in lists:
            if i:
                self.pushHeap(i)
        if not self.data:
            return None
        # 记录头节点
        head = self.popHeap()
        if head.next:
            self.pushHeap(head.next)
        print(self.data)
        node = head
        # 取链表节点串起来，直到堆为空
        while self.data:
            n = self.popHeap()
            print(n.val)
            node.next = n
            if n.next:
                self.pushHeap(n.next)
            node = node.next
        return head
            
            
    # 入堆（小根堆)
    def pushHeap(self,node:ListNode):
        i = len(self.data)
        self.data.append(node)
        while i > 0 and self.data[(i-1)>>1].val > self.data[i].val:
            # 比它的父节点小，向上换
            self.data[(i-1)>>1],self.data[i] = self.data[i],self.data[(i-1)>>1]
            i = (i-1)>>1
    # 出堆
    def popHeap(self)->ListNode:
        res = self.data[0]
        self.data[0] = self.data[-1]
        self.data.pop()
        i = 0
        l = i * 2 + 1
        best = None
        while l < len(self.data):
            # 找出两个子节点中更小的
            if l + 1 < len(self.data) and self.data[l].val > self.data[l+1].val:
                best = l + 1
            else:
                best = l
            # 判断它们当中小的是否比当前值还要小
            if self.data[i].val > self.data[best].val:
                # 交换位置
                self.data[i],self.data[best] = self.data[best],self.data[i]
            i = best
            l = i*2 + 1
        return res
            
            
        