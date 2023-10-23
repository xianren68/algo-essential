from typing import Optional
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        if head == None or head.next == None :
            return True
        slow = head
        quick = head
        while  not quick.next == None and not quick.next.next == None :
            quick = quick.next.next
            slow = slow.next
        rHead = self.reverse(slow.next)
        node = head
        rNode = rHead
        while rNode != None and node != None and node.val == rNode.val:
            rNode = rNode.next
            node = node.next
        return rNode == None or node == None
    def reverse(self,start:Optional[ListNode])->Optional[ListNode]:
        cur = start
        pre = None
        next = None
        while cur != None:
            next = cur.next
            cur.next = pre
            pre = cur
            cur = next
        return pre
            
        