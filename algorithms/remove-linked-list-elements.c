/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* removeElements(struct ListNode* head, int val) {
	struct ListNode* p = head;
	if (head)
		head->next = removeElements(head->next, val);
	else
		return NULL;
	if (val == head->val) {
		head = head->next;
		free(p);
	}
	return head;
}
