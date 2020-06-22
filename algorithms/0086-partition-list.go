package algorithms

func partition(head *ListNode, x int) *ListNode {
	rl := &ListNode{}
	rh := &ListNode{}
	po, pl, ph := head, rl, rh
	for po != nil {
		if po.Val < x {
			pl.Next = po
			pl = pl.Next
		} else {
			ph.Next = po
			ph = ph.Next
		}
		po = po.Next
	}
	pl.Next = rh.Next
    ph.Next = nil
	return rl.Next
}
