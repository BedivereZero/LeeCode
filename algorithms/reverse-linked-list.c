/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* reverseList_iter(struct ListNode* head, struct ListNode* top) {
	if (head == NULL)
		return head;
	if (head->next == NULL) {
		head->next = top;
		return head;
	}
	else
		return reverseList_iter(head->next, head);
}

struct ListNode* reverseList(struct ListNode* head) {
	return reverseList_iter(head->next, head);
}
