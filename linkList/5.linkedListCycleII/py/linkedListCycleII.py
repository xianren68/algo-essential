from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = head
        fast = head
        while True:
            print(1)
            if not(fast and fast.next):
                return None
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                break
        # 快指针回到原点，变成慢指针
        fast = head
        while fast != slow:
            fast = fast.next
            slow = slow.next
        return fast