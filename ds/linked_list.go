package ds

import (
	"fmt"
)

// Node is a node in a linked list
type Node struct {
	next  *Node
	prev  *Node
	value interface{}
}

// LinkedList is a doubly Linked List
type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

// Size returns the current size of the list
func (l *LinkedList) Size() int {
	return l.count
}

// Append adds the value to the end of the l LinkedList
func (l *LinkedList) Append(value interface{}) {
	node := &Node{value: value}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}

	l.count++
}

// Add is an alias for Append
func (l *LinkedList) Add(value interface{}) {
	l.Append(value)
}

// Prepend adds the element at the start of the list
func (l *LinkedList) Prepend(value interface{}) {
	node := &Node{value: value}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		oldHead := l.head
		node.next = oldHead
		oldHead.prev = node
		l.head = node
	}

	l.count++
}

// RemoveFirst removes the head from this list and returns it
func (l *LinkedList) RemoveFirst() interface{} {
	if l.head == nil {
		return -1
	}

	oldHead := l.head
	newHead := l.head.next
	if newHead != nil {
		newHead.prev = nil
	}

	l.head = newHead
	l.count--

	return oldHead.value
}

// Poll is an alias for RemoveFirst
func (l *LinkedList) Poll() interface{} {
	return l.RemoveFirst()
}

// Remove will take the tail of the list, remove it and return its value
func (l *LinkedList) Remove() interface{} {
	if l.tail == nil {
		return nil
	}

	oldTail := l.tail
	fmt.Printf("Old tail is %v\n", oldTail.value)
	if oldTail.prev == nil {
		l.tail = nil
		l.head = nil
	} else {
		newTail := oldTail.prev
		newTail.next = nil
		l.tail = newTail
	}

	l.count--

	return oldTail.value
}

func (n *Node) String() string {
	var prev interface{}
	var next interface{}

	if n.prev != nil {
		prev = n.prev.value
	}

	if n.next != nil {
		next = n.next.value
	}

	return fmt.Sprintf("[%v prev: %v next: %v]", n.value, prev, next)
}

// Empty returns true if the length of the Linked List is Zero
func (l *LinkedList) Empty() bool {
	return l.count == 0
}
