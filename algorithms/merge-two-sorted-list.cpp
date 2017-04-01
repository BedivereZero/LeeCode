struct ListNode {
	int val;
	ListNode * next;
	ListNode(int x): val(x), next(NULL) {}
}

class Solution {
	public:
		ListNode * mergeTwoLists(ListNode * L1, ListNode * L2) {
			if(L1 == NULL && L2 == NULL)
				return NULL;
			if(L1 == NULL && L2 != NULL)
				return L2;
			if(L1 != NULL && L2 == NULL)
				return L1;
			ListNode * P = NULL;
			if(L1 -> val < L2 -> val) {
				P = L1;
				L1 = L1 -> next;
			} else {
				P = L2;
				L2 = L2 -> next;
			}
			P -> next = mergeTwoLists(L1, L2);
			return P;
		}
};
