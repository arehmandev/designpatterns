package main

import "fmt"

// Node -
type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

// List -
type List struct {
	head *Node
	tail *Node
}

// Insert -
func (L *List) Insert(key interface{}) {
	newHead := &Node{
		next: L.head,
		key:  key,
	}

	// Keep track of oldHead
	oldHead := L.head

	// Always insert new head
	L.head = newHead

	if oldHead == nil {
		// If head was originally null, also make it the tail
		L.tail = newHead
	} else {
		// If list head is not empty, add a 'prev' node entry to the oldHead
		oldHead.prev = newHead
	}

}

// Show - iterates from head to tail and prints all keys
func (L *List) Show() {
	list := L.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// Reverse -
func (L *List) Reverse() {
	curr := L.head
	var prev *Node
	L.tail = L.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	L.head = prev
}

func main() {
	l := List{}
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)

	fmt.Printf("head: %v\n", l.head.key)
	fmt.Printf("tail: %v\n", l.tail.key)
	l.Show()

	l.Reverse()
	l.Show()
	fmt.Printf("head: %v\n", l.head.key)
	fmt.Printf("tail: %v\n", l.tail.key)
}
