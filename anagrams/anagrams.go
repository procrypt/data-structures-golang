package main

import (
	"fmt"
	"strings"
	"os"
)

var a string
var b string

func main() {

	c := make(map[string]int)

	fmt.Scan(&a)
	fmt.Scan(&b)

	a = strings.Replace(a, " ", "", -1)
	b = strings.Replace(b, " ", "", -1)


	if len(a) != len(b) {
		fmt.Println("Not an anagram")
		os.Exit(1)
	}

	for _, i := range a {
		_, prs := c[string(i)]
		if prs {
			c[string(i)] += 1
		} else {
			c[string(i)] = 1
		}
	}

	for _, j := range b {
		_, prs := c[string(j)]
		if prs {
			c[string(j)] -= 1
		} else {
			c[string(j)] = 1
		}
	}

	for i := range c {
		if c[i] != 0 {
			fmt.Println("Not Anagram.")
		} else {
			fmt.Println("Anagram.")
		}
	}

}
