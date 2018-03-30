package main

import (
	"errors"
	"fmt"
)

// Create dynamic array struct
type DynamicArray struct {
	size 	  int
	capacity  int
	container []interface{}
}

// Resize() method, resize the dynamic array capacity to twice
func(d *DynamicArray) Resize() {
	// Make a new container of length defined
	newContainer := make([]interface{}, 2*d.capacity)
	// Copy data of previous container to new one
	copy(newContainer, d.container)
	d.container = newContainer
	// Update the capacity to twice
	d.capacity = 2*d.capacity
}

// Get() method, retrieves the value from the given index
func(d *DynamicArray) Get(index int) (interface{}, error){
	if index >= 0 && index < d.size {
		return d.container[index], fmt.Errorf("")
	} else {
		return nil, errors.New("index out of range")
	}
}

// Set() method, sets the value at the given index
func(d *DynamicArray) Set(index int, value interface{}) error {
	if index >= 0 && index < d.size {
		d.container[index] = value
		return nil
	} else {
		return errors.New("index out of range")
	}
}

// Append() method, adds element to the end of the array
func(d *DynamicArray) Append(value interface{}) {
	if d.size <= d.capacity {
		d.Resize()
	}
	d.container[d.size] = value
	d.size += 1
}

// Length() method, gives the length of the array
func (d *DynamicArray) Length() int {
	return len(d.container)
}

func main() {
	array := DynamicArray{size:0, capacity:1,}
	array.Append(1)
	fmt.Println(array.container)
}
