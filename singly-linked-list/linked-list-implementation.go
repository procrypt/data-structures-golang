package main

import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
	size int
}

type Node struct {
	Data interface{}
	Next *Node
}

func New() *LinkedList {
	return &LinkedList{}
}

func(list *LinkedList) Append(item interface{}) {
	newElement := &Node{Data: item}
	if list.size == 0 {
		list.Head = newElement
		list.Tail = newElement
	} else {
		list.Tail.Next = newElement
		list.Tail = newElement
	}
	list.size++
}

func(list *LinkedList) Size() int {
	return list.size
}

func(list *LinkedList) IsEmpty() bool {
	if list.size == 0 {
		return true
	} else {
		return false
	}
}

func(list *LinkedList) String() {
	head := list.Head
	for {
		if head == nil {
			break
		} else {
			fmt.Println(head.Data)
			head = head.Next
		}
	}
}

func(list *LinkedList) head() *Node {
	return list.Head
}

func(list *LinkedList) indexOf(item interface{}) int {
	head := list.Head
	counter := 0
	for {
		if head.Data == item {
			return counter
		}
		head = head.Next
		counter++
	}
}


func main() {
	a := New()
	for i := 0; i < 5; i++ {
		a.Append(i)
	}

	//a.String()
	fmt.Println(a.indexOf(4))
}
