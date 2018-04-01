package main

import "fmt"

func main() {

	a := []int{1,2,3,4,5,6,7,8}
	b := []int{1,3,4,5,6,8,2}

	c := make(map[int]int)

	for _,i := range a {
			c[i] += 1
	}

	for _, j := range b {
		c[j] -= 1
	}

	for k,_ := range c {
		if c[k] != 0 {
			fmt.Println(k)
		}
	}
}
