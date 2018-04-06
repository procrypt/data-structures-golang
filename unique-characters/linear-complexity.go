package main

import "fmt"

func uniqueCharater(s string) bool {
	c := make(map[string]int)
	if len(s) == 1 {
		return true
	}

	for i := range s {
		_,ok := c[string(s[i])]
		if ok {
				return false
		} else {
			c[string(s[i])] += 1
		}
	}
	return true
}

func main() {
	a := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(uniqueCharater(a))
}
