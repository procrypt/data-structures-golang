package main

import "fmt"

func Factorial(i int) int {
	if i == 0 {
		return 1
	} else {
		return i * Factorial(i-1)
	}
}
func main() {
	fmt.Println(Factorial(10))
}
