package algorithms

type (
	node struct {
		val  interface{}
		next *node
	}
	AscedingList struct {
		head *node
	}
	MinStack struct {
		top   *node
		min   *node
		depth int
		asc   *AscedingList
	}
)

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{asc: &AscedingList{}}
}

func (this *MinStack) Push(x int) {
	n := &node{val: x}
	if this.depth == 0 {
		this.top = n
		this.min = n
	} else {
		n.next = this.top
		this.top = n
	}
	this.depth++
	if this.min == nil || n.val.(int) < this.min.val.(int) {
		this.min = n
	}
	this.asc.Insert(n.val)
}

func (this *MinStack) Pop() {
	if this.depth == 0 {
		return
	}

	n := this.top
	if this.depth == 1 {
		this.top = nil
	} else {
		this.top = this.top.next
	}
	this.depth--

	this.asc.Delete(n.val)
	this.min = this.asc.head
}

func (this *MinStack) Top() int {
	return this.top.val.(int)
}

func (this *MinStack) GetMin() int {
	return this.min.val.(int)
}

/** asceding list */
func (l *AscedingList) Insert(element interface{}) {
	n := &node{val: element}
	if l.head == nil {
		l.head = n
		return
	}

	if n.val.(int) < l.head.val.(int) {
		n.next = l.head
		l.head = n
		return
	}

	cur := l.head
	for cur.next != nil {
		if n.val.(int) < cur.next.val.(int) {
			n.next = cur.next
			cur.next = n
			return
		}
		cur = cur.next
	}
	cur.next = n
}

func (l *AscedingList) Delete(element interface{}) {
	if l.head.val.(int) == element.(int) {
		l.head = l.head.next
		return
	}

	cur := l.head
	for cur.next != nil {
		if cur.next.val.(int) == element.(int) {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
