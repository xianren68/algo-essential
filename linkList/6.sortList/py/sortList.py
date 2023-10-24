from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
        
class Solution:
    start:Optional[ListNode] = None
    end:Optional[ListNode] = None
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        l = 0 
        node = head
        while node:
            l+=1
            node = node.next
        i = 1
        next,lastTeamEnd = None,None
        
        while i < l :
            l1 = head
            r1 = self.findEnd(l1,i)
            l2 = r1.next
            r2 = self.findEnd(l2,i)
            next = r2.next
            r1.next = None
            r2.next = None
            self.merge(l1,r1,l2,r2)
            head = self.start
            lastTeamEnd = self.end
            while next:
                l1 = next
                r1 = self.findEnd(l1,i)
                l2 = r1.next
                if not l2:
                    lastTeamEnd.next = l1
                    break
                r2 = self.findEnd(l2,i)
                next = r2.next
                r1.next = None
                r2.next = None
                self.merge(l1,r1,l2,r2)
                lastTeamEnd.next = self.start
                lastTeamEnd = self.end
            i <<= 1
        return head
            
    def findEnd(self,start:Optional[ListNode],k:int):
        while start.next and k > 1:
            start = start.next
            k -= 1
        return start

    def merge(self,l1:Optional[ListNode],r1:Optional[ListNode],l2:Optional[ListNode],r2:Optional[ListNode]):
        start,end = None,None
        if l1.val < l2.val:
            start = l1
            l1 = l1.next
        else:
            start = l2
            l2 = l2.next
        node = start
        while l1 and l2 :
           if l1.val < l2.val:
               node.next = l1
               l1 = l1.next
               node = node.next
           else:
               node.next = l2
               l2 = l2.next
               node = node.next
        if l1:
            node.next = l1
            end = r1
        else:
            node.next = l2
            end = r2
        self.start = start
        self.end = end