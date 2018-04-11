package singly_linked_list

import "fmt"

type element struct {
	 value interface{}
	 next  *element
}

type List struct {
	 first *element
	 last  *element
	 size   int
}

func New() *List {
	return &List{}
}


// Add elements to the end of the list
func(list *List) Add(values ...interface{}) {
	for _,value := range values {
		newElement := &element{value: value} // Initialize a new list
		if list.size == 0 {                  // If the size of the list is 0, means a new list, first and last element are same
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement      // Next to the last element
			list.last = newElement 	         // New Last Element
		}
		list.size++						     // Increment the list size
	}

}


func main() {

	a := New()
	a.Add("Abhishek", "Pratap", "Singh", "New York City", "New York")
	a.Add("USA")

	fmt.Println(a.first.value)
	fmt.Println(a.first.next.value)
	fmt.Println(a.first.next.next.value)
	fmt.Println(a.first.next.next.next.value)
	fmt.Println(a.first.next.next.next.next.value)
	fmt.Println(a.first.next.next.next.next.next.value)

}