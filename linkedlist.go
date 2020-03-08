package golist

import (
	"fmt"
)

// Node type is an element of a linked list that contains references to the previous
// (prev) and next (next) elements, which contains value (val).
type Node struct {
	prev *Node
	next *Node
	val  interface{}
}

// List is a struct comprising of a reference to the head node and tail node of a
// List type
type List struct {
	head *Node
	tail *Node
}

// Insert placed a new Node type struct at the head of a list and references the
// Node as it's next Node.
func (l *List) Insert(val interface{}) {
	lst := &Node{
		next: l.head,
		val:  val,
	}
	if l.head != nil {
		l.head.prev = lst
	}
	l.head = lst

	li := l.head
	for li.next != nil {
		li = li.next
	}
	l.tail = li
}

// Append take a value and places a new Node struct as the tail and previous tail
// as the that Node's prev parameter
func (l *List) Append(val interface{}) {
	newTail := &Node{
		prev: l.tail,
		val:  val,
	}
	l.tail.next = newTail
	l.tail = newTail
}

// FindNode iterates from head to tail of a list looking for the pass value (val)
func (l *List) FindNode(val interface{}) (bool, *Node) {
	lst := l.head
	for lst != nil {
		if lst.val == val {
			return true, lst
		}
		lst = lst.next
	}
	return false, nil
}

// RemNode uses FindNode to vind the first element with value val (if there),
// then checks if the Node containing value val is the head or not... removing as needed
func (l *List) RemNode(val interface{}) bool {

	isThere, current := l.FindNode(val)
	if isThere && (current != l.head) {
		newHead := current.prev
		newHead.next = current.next
		current = nil
		return true
	} else if isThere && (current == l.head) {
		newHead := current.next
		newHead.prev = nil
		current = nil
		l.head = newHead
		return true

	}
	return false

}

// Display prints  to std.out an entire list
func (l *List) Display() {
	lst := l.head
	for lst != nil {
		if lst.next != nil {
			fmt.Printf("%+v -> ", lst.val)
		} else {
			fmt.Printf("%+v", lst.val)
		}
		lst = lst.next
	}
	fmt.Println()
}

// Display here takes a Node struct and prints the list after that Node
func Display(lst *Node) {
	for lst != nil {
		if lst.next != nil {
			fmt.Printf("%v -> ", lst.val)
		} else {
			fmt.Printf("%v", lst.val)
		}
		lst = lst.next
	}
	fmt.Println()
}

// ShowBackwards takes a Node struct and prints the list before that Node
func ShowBackwards(lst *Node) {
	for lst != nil {
		fmt.Printf("%v <-", lst.val)
		lst = lst.prev
	}
	fmt.Println()
}

// Reverse completely reverses a List's Node objects links
func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	//Display(l.head)

}
