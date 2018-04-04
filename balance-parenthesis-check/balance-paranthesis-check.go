package main

import "fmt"

func check(s string) bool {
	c := make(map[byte]rune)
	if len(s)%2 == 0 {
		for i := range s {
			_, ok := c[byte(s[i])]
			if ok {
				c[byte(s[i])] += 1
			} else {
				c[byte(s[i])] = 1
			}
		}
	} else {
		return false
	}

	if c[40] == c[41] && c[91] == c[93] && c[123] == c[125]  {
		return true
	} else {
		return false
	}
}

func main() {
	s := "()(){][}"
	fmt.Println(check(s))
}
