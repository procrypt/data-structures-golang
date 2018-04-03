package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "Hi John, are you ready to go?"
	b := strings.Split(strings.TrimSpace(a), " ")
	j := 1

	if len(b)%2 == 0 {
		c := len(b) / 2
		for i := 0; i < c; i++ {
			temp := b[i]
			b[i] = b[len(b)-j]
			b[len(b)-j] = temp
			j++
		}
	} else {
		c := (len(b) / 2) + 1
		for i := 0; i < c; i++ {
			temp := b[i]
			b[i] = b[len(b)-j]
			b[len(b)-j] = temp
			j++
		}
	}
	fmt.Println(b)
}
