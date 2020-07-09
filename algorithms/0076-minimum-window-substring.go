package algorithms

type (
	queue struct {
		head, tail *node
		length     int
	}
	node struct {
		value interface{}
		next  *node
	}
	count struct {
		actual, expect int
	}
	SliceWindow struct {
		head, tail     int
		source, target string
		indexQueue     queue
		record         map[byte]*count
		targetCount    int
	}
)

func (q *queue) enqueue(element interface{}) {
	n := &node{value: element}
	if q.length == 0 {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.length++
}

func (q *queue) dequeue() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.head
	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.length--
	return n.value
}

func genSlideWindow(s, t string) *SliceWindow {
	w := SliceWindow{source: s, target: t}
	w.record = make(map[byte]*count)
	for i := range t {
		if w.checkInTarget(t[i]) {
			w.record[t[i]].expect++
		} else {
			w.record[t[i]] = &count{expect: 1}
		}
	}
	return &w
}

func (w *SliceWindow) Len() int {
	return w.tail - w.head
}

func (w *SliceWindow) Expand() {
	for ; !w.checkContainsALL() && w.tail < len(w.source); w.tail++ {
		if w.checkInTarget(w.source[w.tail]) {
			w.indexQueue.enqueue(w.tail)
			w.record[w.source[w.tail]].actual++
			if w.record[w.source[w.tail]].actual == w.record[w.source[w.tail]].expect {
				w.targetCount++
			}
		}
	}
}

func (w *SliceWindow) Reduce() {
	if w.indexQueue.length < len(w.target) || !w.checkContainsALL() {
		return
	}
	for w.checkContainsALL() {
		w.head, _ = w.indexQueue.dequeue().(int)
		if w.record[w.source[w.head]].actual == w.record[w.source[w.head]].expect {
			w.targetCount--
		}
		w.record[w.source[w.head]].actual--
	}
	w.head++
}

func (w *SliceWindow) checkContainsALL() bool {
	return w.targetCount == len(w.record)
}

func (w *SliceWindow) checkInTarget(b byte) bool {
	_, ok := w.record[b]
	return ok
}

func minWindow(s string, t string) string {
	var substring string
	var exists bool
	w := genSlideWindow(s, t)
	for w.tail < len(s) {
		w.Expand()
		w.Reduce()
		if w.head > 0 && w.targetCount+1 == len(w.record) && w.record[s[w.head-1]].actual+1 == w.record[s[w.head-1]].expect {
			if exists {
				if w.Len()+1 < len(substring) {
					substring = s[w.head-1 : w.tail]
				}
			} else {
				substring = s[w.head-1 : w.tail]
				exists = true
			}
		}
	}
	return substring
}

func checkAllExist(expect, actual map[byte]int) bool {
	for k, v := range expect {
		if actual[k] < v {
			return false
		}
	}
	return true
}
