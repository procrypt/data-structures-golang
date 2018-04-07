package main

import "fmt"

type Queue struct {
	container []interface{}
}

func(q *Queue) Enqueue(item interface{}) interface{} {
	a := append(q.container, item)
	q.container = a
	return q.container
}

func(q *Queue) Dequeue() interface{} {
	r := q.container[0]
	a := q.container[1:]
	q.container = a
	return r
}

func(q *Queue) isEmpty() bool {
	if len(q.container) == 0 {
		return true
	} else {
		return false
	}
}

func(q *Queue) Size() int {
	return len(q.container)
}

func main() {

	a := Queue{}
	fmt.Println(a.isEmpty())
	a.Enqueue("Abhishek")
	a.Enqueue("Pratap")
	a.Enqueue("Singh")
	fmt.Println(a.Size())
	fmt.Println(a.isEmpty())
	fmt.Println(a.container)
	fmt.Println(a.Dequeue())
	fmt.Println(a.Size())
	fmt.Println(a.container)
	fmt.Println(a.Dequeue())
	fmt.Println(a.Size())
	fmt.Println(a.container)
	fmt.Println(a.Dequeue())
	fmt.Println(a.Size())
	fmt.Println(a.container)
	fmt.Println(a.Size())
	fmt.Println(a.isEmpty())

}
