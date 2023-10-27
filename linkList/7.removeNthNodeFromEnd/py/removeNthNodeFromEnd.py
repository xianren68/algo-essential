from typing import Optional
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        if not head:
            return None
        # 定义前后两个指针
        before, after = head, head
        # 前指针先走n步
        i = 0
        while before and i < n:
            before = before.next
            i += 1
        if not before:
            if i < n:
                # 长度不足n,原样返回
                return head
            else:
                # 长度恰好为n,删除头节点
                return head.next
        # 指向前一个节点，用于删除节点
        pre: Optional[ListNode] = None
        # 前后指针一起走，前面走到末尾，后面正好到倒数第n个
        while before:
            pre = after
            after = after.next
            before = before.next
        pre.next = after.next
        return head