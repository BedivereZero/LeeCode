# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        l3 = None
        n1, n2 = l1, l2
        pre = l3
        adv = 0
        while n1 or n2 or adv:
            v1, v2 = n1.val if n1 else 0, n2.val if n2 else 0
            ret = v1 + v2 + adv
            adv = ret // 10
            now = ListNode(x=ret % 10)
            if pre:
                pre.next = now
            else:
                l3 = now
            n1, n2, pre = n1.next if n1 else None, n2.next if n2 else None, now
        return l3
