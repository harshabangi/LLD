package main

import "fmt"

type Node struct {
	Key   int
	Value int
	prev  *Node
	next  *Node
}

func NewNode(key, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

type DoublyLinkedList struct {
	len  int
	head *Node
	tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (d *DoublyLinkedList) Len() int {
	return d.len
}

func (d *DoublyLinkedList) Front() *Node {
	return d.head
}

func (d *DoublyLinkedList) Back() *Node {
	return d.tail
}

func (d *DoublyLinkedList) MoveToFront(key int) {
	curr := d.head
	for curr != nil {
		if curr.Key == key {

			if curr.Key == d.head.Key {
				return
			}

			if curr.prev != nil {
				curr.prev.next = curr.next
			}

			if curr.next != nil {
				curr.next.prev = curr.prev
			} else {
				d.tail = curr.prev
			}

			curr.prev = nil
			curr.next = d.head
			d.head.prev = curr
			d.head = curr
		}
		curr = curr.next
	}
}

func (d *DoublyLinkedList) PushFront(newNode *Node) {
	d.len++

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	newNode.next = d.head
	d.head.prev = newNode
	d.head = newNode
}

func (d *DoublyLinkedList) PushEnd(newNode *Node) {
	d.len++

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	newNode.prev = d.tail
	d.tail.next = newNode
	d.tail = d.tail.next
}

func (d *DoublyLinkedList) RemoveFromEnd() {
	if d.head == nil {
		return
	}
	d.len--

	if d.head == d.tail {
		d.head = nil
		d.tail = nil
		return
	}

	pr := d.tail
	d.tail = d.tail.prev
	d.tail.next = nil
	pr.prev = nil
}

func (d *DoublyLinkedList) Iterate() {
	head := d.head

	for head != nil {
		fmt.Printf("[%d,%d]", head.Key, head.Value)
		fmt.Printf(" -> ")
		head = head.next
	}
	fmt.Printf(" nil\n")
}

type LRUCache struct {
	dll      DoublyLinkedList
	store    map[int]*Node
	capacity int
	length   int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		dll:      DoublyLinkedList{},
		store:    make(map[int]*Node),
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	v, ok := l.store[key]
	if !ok {
		return -1
	}
	l.dll.MoveToFront(key)
	return v.Value
}

func (l *LRUCache) Put(key, value int) {
	v, present := l.store[key]

	switch {
	case present && l.length <= l.capacity:
		v.Value = value
		l.dll.MoveToFront(key)
		l.store[key] = l.dll.head

	case !present && l.length < l.capacity:
		newNode := NewNode(key, value)
		l.dll.PushFront(newNode)
		l.store[key] = newNode
		l.length++

	case !present && l.length == l.capacity:
		back := l.dll.Back()
		l.dll.RemoveFromEnd()
		delete(l.store, back.Key)

		newNode := NewNode(key, value)
		l.dll.PushFront(newNode)
		l.store[key] = newNode
	}
	return
}
