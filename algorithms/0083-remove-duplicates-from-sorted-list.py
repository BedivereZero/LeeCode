# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        now = head
        while now is not None:
            while now.next is not None and now.val == now.next.val:
                now.next = now.next.next
            now = now.next
        return head
