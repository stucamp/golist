package linkedlist

import (
	"fmt"
)

type node struct {
	prev *node
	next *node
	key  interface{}
}

type list struct {
	head *node
	tail *node
}

func (l *list) insert(key interface{}) {
	lst := &node{
		next: l.head,
		key:  key,
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

func (l *list) append(key interface{}) {
	newTail := &node{
		prev: l.tail,
		key:  key,
	}
	l.tail.next = newTail
	l.tail = newTail
}

func (l *list) findNode(key interface{}) (bool, *node) {
	lst := l.head
	for lst != nil {
		if lst.key == key {
			return true, lst
		}
		lst = lst.next
	}
	return false, nil
}

func (l *list) remNode(key interface{}) bool {

	isThere, current := l.findNode(key)
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

func (l *list) display() {
	lst := l.head
	for lst != nil {
		if lst.next != nil {
			fmt.Printf("%+v -> ", lst.key)
		} else {
			fmt.Printf("%+v", lst.key)
		}
		lst = lst.next
	}
	fmt.Println()
}

func display(lst *node) {
	for lst != nil {
		if lst.next != nil {
			fmt.Printf("%v -> ", lst.key)
		} else {
			fmt.Printf("%v", lst.key)
		}
		lst = lst.next
	}
	fmt.Println()
}

func showBackwards(lst *node) {
	for lst != nil {
		fmt.Printf("%v <-", lst.key)
		lst = lst.prev
	}
	fmt.Println()
}

func (l *list) reverse() {
	curr := l.head
	var prev *node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	display(l.head)

}
