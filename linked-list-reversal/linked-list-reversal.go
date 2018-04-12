package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	nextNode *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func New() *LinkedList {
	return &LinkedList{}
}

func(list *LinkedList) Add(values ...interface{}) {
	for _,value := range values {
		newElement := &Node{value: value}
		if list.size == 0 {
			list.head = newElement
			list.tail = newElement
		} else {
			list.tail.nextNode  = newElement
			list.tail = newElement
		}
		list.size++
	}
}

func Reverse(l *LinkedList) {
	current := l.head			// set the current element to be the head of the list
	var prev  *Node				// variable that holds the previous link
	var next  *Node             // variable that holds the next link

	for current != nil {
		next = current.nextNode  // next holds the link to the next node of the current element
		current.nextNode = prev  // set current.nextNode link to the previous node link
		prev = current           // set previous node as current node
		current = next           // set the current node as next node
	}
	l.head = prev                // Make sure the new head which is the previous tail now point to the previous
}                                // element to the tail not NULL,

func Show(l *LinkedList) {
	for l.head != nil {
		fmt.Println(l.head.value)
		l.head = l.head.nextNode
	}
}

func main() {

	a := New()
	a.Add(1,2,3,4,5)
	Reverse(a)
	Show(a)
}