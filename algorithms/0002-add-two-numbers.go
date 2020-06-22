package algorithms

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode;
	var pr *ListNode;
	var ad int = 0;
	for (l1 != nil) || (l2 !=  nil) || (ad != 0) {
		var v1, v2 int;
		if l1 != nil {
			v1 = l1.Val;
		} else {
			v1 = 0;
		}
		if l2 != nil {
			v2 = l2.Val;
		} else {
			v2 = 0;
		}
		var re int = v1 + v2 + ad;
		var ne ListNode;
		ad = re / 10;
		ne.Val = re % 10;
		ne.Next = nil;
		if pr != nil {
			pr.Next = &ne;
		} else {
			l3 = &ne;
			pr = l3;
		}
		if l1 != nil {
			l1 = l1.Next;
		} else {
			l1 = nil;
		}
		if l2 != nil {
			l2 = l2.Next;
		} else {
			l2 = nil;
		}
		pr = &ne;
	}
	return l3;
}
