package main

import "fmt"

type Stack struct {
	container []interface{}
}

func(s *Stack) isEmpty() bool {
	if len(s.container) == 0 {
		return true
	} else {
		return false
	}
}

func(s *Stack) pop() interface{} {
	a := s.container[:len(s.container)-1]
	s.container = a
	return a
}

func(s *Stack) peek() interface{} {
	return s.container[len(s.container)-1:]
}

func (s *Stack) push(item interface{}) interface{} {
	s.container = append(s.container, item)
	return s.container
}

func (s *Stack) size() int {
	return len(s.container)
}

func main() {

	a := Stack{}
	fmt.Println(a.isEmpty())
	a.push("Abhishek")
	a.push("Pratap")
	a.push("Singh")
	a.push("New York City")
	a.push("Jaipur")
	a.push("Hello")
	fmt.Println(a.container)
	fmt.Println(a.size())
	fmt.Println(a.pop())
	fmt.Println(a.size())
}
