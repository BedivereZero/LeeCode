# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1 is None:
            return l2
        if l2 is None:
            return l1
        head = ListNode(0)
        l3 = head
        while not (l1 is None and l2 is None):
            if l1 is None:
                l3.next = l2
                l2 = l2.next
            elif l2 is None:
                l3.next = l1
                l1 = l1.next
            elif l1.val < l2.val:
                l3.next = l1
                l1 = l1.next
            else:
                l3.next = l2
                l2 = l2.next
            l3 = l3.next
        return head.next
