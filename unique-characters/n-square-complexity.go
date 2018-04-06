package main

import "fmt"

func unique(s string) bool {
	c := make(map[string]int)

	for i := range s {
		c[string(s[i])] += 1
	}

	for k,_ := range c {
		if c[k] > 1 {
			return false
			break
		}
	}
	return true
}

func main() {
	a := "abcdefghij"
	fmt.Println(unique(a))
}
