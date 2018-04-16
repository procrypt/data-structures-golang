package main

import (
	"fmt"
)

func Reverse(s string) string {
	if len(s) == 1 {
		return s
	} else {
		return Reverse(s[1:]) + string(s[0])
	}
}

func main() {
	fmt.Println(Reverse("Hello"))
}
