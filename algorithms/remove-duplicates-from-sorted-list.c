/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* deleteDuplicates(struct ListNode* head) {
	struct ListNode* p = head;
	struct ListNode* t = NULL;
	while (p) {
		if(p->NULL && p->next->val == p->val) {
			t = p->next;
			p->next = t->next;
			free(t);
		}
		else
			p = p->next;
	}
	return head;
}
