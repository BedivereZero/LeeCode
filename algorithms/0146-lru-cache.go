package algorithms

import (
	"fmt"
	"log"
)

type (
	node struct {
		key       interface{}
		val       interface{}
		pre, next *node
	}
	LRUCache struct {
		nodeFromKey map[int]*node
		head, tail  *node
		used        int
		capacity    int
	}
)

func Constructor(capacity int) LRUCache {
	return LRUCache{
		nodeFromKey: map[int]*node{},
		capacity:    capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	log.Println("Get:", key)
	n, ok := this.nodeFromKey[key]
	if !ok {
		return -1
	}
	if this.head == n {
		return n.val.(int)
	}

	n.pre.next = n.next
	if this.tail == n {
		this.tail = n.pre
	} else {
		n.next.pre = n.pre
	}

	n.pre = nil
	n.next = this.head
	this.head.pre = n
	this.head = n
	return n.val.(int)
}

func (this *LRUCache) Put(key int, value int) {
	log.Println("Put:", key, value)
	if n, ok := this.nodeFromKey[key]; ok {
		n.val = value
		if this.head == n {
			return
		}

		n.pre.next = n.next

		if this.tail == n {
			this.tail = n.pre
		} else {
			n.next.pre = n.pre
		}

		n.pre = nil
		n.next = this.head
		this.head.pre = n
		this.head = n
		return
	}

	n := &node{key: key, val: value}
	this.nodeFromKey[key] = n

	if this.used == 0 {
		this.head, this.tail = n, n
		this.used = 1
		return
	}

	n.next = this.head

	this.head.pre = n
	this.head = n

	if this.used < this.capacity {
		this.used++
		return
	}

	if this.tail.pre != nil {
		delete(this.nodeFromKey, this.tail.key.(int))
		this.tail = this.tail.pre
		this.tail.next = nil
	}
}

func (this *LRUCache) Debug(prefix string) {
	log.Printf("%s: %d/%d, map=%s, forward=%s, reverse=%s", prefix, this.used, this.capacity, printMap(this.nodeFromKey), printListForward(this.head), printListReverse(this.tail))
}

func printMap(m map[int]*node) string {
	s := "["
	for _, n := range m {
		s += fmt.Sprintf("%d:%d, ", n.key.(int), n.val.(int))
	}
	return s + "]"
}

func printListForward(head *node) string {
	s := "["
	for n := head; n != nil; n = n.next {
		s += fmt.Sprintf("%d:%d, ", n.key.(int), n.val.(int))
	}
	return s + "]"
}

func printListReverse(tail *node) string {
	s := "["
	for n := tail; n != nil; n = n.pre {
		s += fmt.Sprintf("%d:%d, ", n.key.(int), n.val.(int))
	}
	return s + "]"
}
