package main

import "fmt"

var cache map[int]int

func Fibonacci(i int) int {
	if cache == nil {
		cache = make(map[int]int)
	} else {
		_, ok := cache[i]
		if ok {
			return cache[i]
		}
	}
	if i == 1 {
		return 1
	} else if i == 2 {
		return 1
	} else {
		cache[i] = Fibonacci(i-1) + Fibonacci(i-2)
		return cache[i]
	}
}

func main() {
	for i:=1; i<50; i++ {
		fmt.Println(Fibonacci(i))
	}
}
