package main

import (
	"fmt"
)

type element struct {
	value interface{}
	next *element
}

type List struct {
	first *element
	last  *element
	size int
}

func(list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

func New() *List {
	return &List{}
}

func check(list *List) {

	marker1 := list
	marker2 := list

	for marker2 != nil && marker2.first.next != nil {

		marker1 = marker1.first.next
		marker2 = marker2.next.next

		if marker1 == marker2 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
func main() {

		a := New()
		a.Add("1")

		check(a)

		a.Add("2", "1")
		check(a)
		fmt.Println(a.first.value)
		fmt.Println(a.first.next.value)
		fmt.Println(a.first.next.next)

}
