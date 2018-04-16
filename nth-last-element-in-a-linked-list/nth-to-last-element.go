package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	nextNode *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func New() *List {
	return &List{}
}

func(list *List) Add(values ...interface{}) {
	for _, value :=  range values {
		newElement := &Node{value: value}
		if list.size == 0 {
			list.head = newElement
			list.tail = newElement
		} else {
			list.tail.nextNode = newElement
			list.tail = newElement
		}
		list.size++
	}
}


func Reverse(l *List) {
	current := l.head
	var prev *Node
	var next *Node

	for current != nil {
		next = current.nextNode
		current.nextNode = prev
		prev = current
		current = next
	}
	l.head = prev
}

func Show(l *List) {
	for l.head != nil {
		fmt.Println(l.head.value)
		l.head = l.head.nextNode
	}
}

func GetNthElement(i int, l *List) {
	Reverse(l)                        // Reverse the linked list for ease in looping
	count := 1
	for l.head != nil {               // Loop over the list
		if count == i {               // If count is equal to i, print that element
			fmt.Println(l.head.value)
		} else {
			l.head = l.head.nextNode
		}
		count++
	}
}

func main() {
	a := New()
	a.Add(1,2,3,4,5,6,7,8,9)
	GetNthElement(3, a)
}
