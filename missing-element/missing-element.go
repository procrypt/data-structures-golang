package main

import (
	"sort"
	"fmt"
)

func main() {
	a := []int{1,2,3,4,5,6,7,8}
	b := []int{1,3,4,5,6,7,2}

	sort.Ints(a)
	sort.Ints(b)

	for i := range b {
		if a[i] != b[i] {
			fmt.Println(a[i])
			break
		}
	}
		fmt.Println(a[len(a)-1])
}
