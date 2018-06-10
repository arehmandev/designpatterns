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
		// If oldHead was originally null, also make it the tail
		L.tail = newHead
	} else {
		// If list oldHead is not empty, add a 'prev' node entry to the oldHead
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
	// Start at tail
	currentNode := L.tail
	// currentNode.next = currentNode.prev

	for currentNode.prev != nil {
		prevNode := currentNode.prev
		nextNode := currentNode.next
		currentNode.next = prevNode
		currentNode.prev = nextNode
	}

	L.tail.next = nil

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
