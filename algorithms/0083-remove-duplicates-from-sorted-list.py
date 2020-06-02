# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        now = head
        while now and now.next:
            if now.val == now.next:
                now.next = now.next.next
            else:
                now = now.next
        return head
