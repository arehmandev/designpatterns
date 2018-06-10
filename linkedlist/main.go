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
	fmt.Println("[FORWARD PRINT] Printing each node, iteration through .next to each next node")
	list := L.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// ShowBackwards -
func (L *List) ShowBackwards() {
	fmt.Println("[REVERSE PRINT] Printing each node, iteration through .prev to each previous node")
	list := L.tail
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

// 1 -> 2 -> 3 -> 4 -> nil
// nil <- 1 <- 2 <- 3 <- 4

// Reverse -
func (L *List) Reverse() {
	fmt.Println("[REVERSING LINKED LIST]")
	curr := L.head
	L.tail = L.head

	var prevNode *Node

	for curr != nil {
		nextNode := curr.next
		curr.next = prevNode
		curr.prev = nextNode
		prevNode = curr
		curr = nextNode
	}
	L.head = prevNode
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
	l.ShowBackwards()

	fmt.Println("===========================================")
	l.Reverse()
	fmt.Printf("head: %v\n", l.head.key)
	fmt.Printf("tail: %v\n", l.tail.key)
	l.Show()
	l.ShowBackwards()

}

/**
Start from head Node

store NextNode

currentNode.prev = currentNode.next
currentNode.next = currentNode.prev

loop through nodes, until reaching currentNode.Next = nil


**/
