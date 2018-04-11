package main

import "fmt"

type element struct {
	value interface{}
	next  *element
	prev  *element
}

type List struct {
	first *element
	last  *element
	size  int
}

func New() *List {
	return &List{}
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			newElement.prev = list.last
			list.last = newElement
		}
		list.size++
	}
}

func main() {

	a := New()
	a.Add("Abhishek", "Pratap", "Singh")

	fmt.Println("++Forward++")

	fmt.Println(a.first.value)
	fmt.Println(a.first.next.value)
	fmt.Println(a.first.next.next.value)

	fmt.Println("++Backward++")
	fmt.Println(a.first.next.value)
	fmt.Println(a.first.next.prev.value)
	fmt.Println(a.first.next.next.value)
	fmt.Println(a.first.next.next.prev.value)
}
