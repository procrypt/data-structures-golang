package main

import "fmt"

var (
	max = 0
	count = 0
)

func main() {
	a := []int{1,2,-1,3,4,10,10,-10,-1}
	for _, i := range a {
		count += i
		if count > max {
			max = count
		}
	}
	fmt.Println(max)
}
