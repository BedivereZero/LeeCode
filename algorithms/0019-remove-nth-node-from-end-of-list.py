# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        a = head
        for _ in range(n):
            a = a.next
        if a is None:
            return head.next
        b = head
        while a.next is not None:
            a, b = a.next, b.next
        b.next = b.next.next
        return head
