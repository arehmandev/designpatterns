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
	newNode := &Node{
		next: L.head,
		key:  key,
	}

	// If list head is not empty, add a 'prev' node entry to the head
	if L.head != nil {
		L.head.prev = newNode
		L.tail = newNode
	}

	L.head = newNode

	// If list tail is empty, then loop through from head next, and set the last node as the tail
	if L.tail == nil {
		l := L.head
		for l.next != nil {
			l = l.next
		}
		L.tail = l
	}
}

// Show -
func (L *List) Show() {
	list := L.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// Show -
func Show(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// ShowBackwards -
func ShowBackwards(list *Node) {
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
	Show(L.head)
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

	fmt.Printf("head: %v\n", l.head.key)
	fmt.Printf("tail: %v\n", l.tail.key)
}
