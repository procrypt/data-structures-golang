package main

import "fmt"

type Deque struct {
	container []interface{}
}

func(d *Deque) addFront(item interface{}) interface{} {
	a := Deque{}.container
	a = append(a, item)
	b := append(a, d.container...)
	d.container = b
	return d.container
}

func(d *Deque) addRear(item interface{}) interface{} {
	a := append(d.container, item)
	d.container = a
	return d.container
}

func(d *Deque) removeFront(item interface{}) interface{} {
	if d.Size(d.container) >= 1 {
		a := d.container[0]
		b := d.container[1:]
		d.container = b
		return a
	} else {
		return fmt.Sprint("Empty container!")
	}
}

func(d *Deque) removeRear(item interface{}) interface{} {
	if d.Size(d.container) >= 1 {
		a := d.container[len(d.container)-1]
		b := d.container[:len(d.container)-1]
		d.container = b
		return a
	} else {
		return fmt.Sprint("Empty container!")
	}
}

func(d *Deque) isEmpty(item interface{}) bool {
	if len(d.container) == 0 {
		return true
	} else {
		return false
	}
}

func(d *Deque) Size(item interface{}) int {
	return len(d.container)
}

func main() {

	a := Deque{}
	fmt.Println(a.isEmpty(a))
	fmt.Println(a.Size(a))
	fmt.Println(a.addFront("Abhishek"))
	fmt.Println(a.addRear("Pratap"))
	fmt.Println(a.addRear("Singh"))
	fmt.Println(a.isEmpty(a))
	fmt.Println(a.removeFront(a))
	fmt.Println(a.addRear("Hello"))
	fmt.Println(a.removeRear(a))
	fmt.Println(a.container)
	fmt.Println(a.removeFront(a))
	fmt.Println(a.removeRear(a))
	fmt.Println(a.removeFront(a))
	fmt.Println(a.removeRear(a))
}
