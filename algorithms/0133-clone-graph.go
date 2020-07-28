package algorithms

type (
	node struct {
		Val  interface{}
		Next *node
	}
	Queue struct {
		start, end *node
		length     int
	}
)

func (q *Queue) Enqueue(value interface{}) {
	n := &node{Val: value}
	if q.length == 0 {
		q.start, q.end = n, n
	} else {
		q.end.Next = n
		q.end = q.end.Next
	}
	q.length++
}

func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}

	n := q.start
	if q.length == 1 {
		q.start, q.end = nil, nil
	} else {
		q.start = q.start.Next
	}
	q.length--
	return n.Val
}

func (q *Queue) Len() int {
	return q.length
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	newNode := &Node{Val: node.Val}
	q := &Queue{}
	q.Enqueue(node)
	mark := map[*Node]*Node{node: newNode}
	for q.Len() > 0 {
		old := q.Dequeue().(*Node)
		new := mark[old]
		for _, neighbor := range old.Neighbors {
			n, ok := mark[neighbor]
			if !ok {
				n = &Node{Val: neighbor.Val}
				mark[neighbor] = n
				q.Enqueue(neighbor)
			}
			new.Neighbors = append(new.Neighbors, mark[neighbor])
		}
	}
	return newNode
}
