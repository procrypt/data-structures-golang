package main

import (
	"fmt"
	"strconv"
)

func compress(s string) string {

	r := ""

	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		str := strconv.Itoa(1)
		return ""+str
	}
	cnt := 1

	for i:=1; i<len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			ctr := strconv.Itoa(cnt)
			r = r+string(s[i-1])+ctr
			cnt = 1
		}
	}
	fmt.Println(cnt)
	return r
}

func main() {

	a := "AAABBCCCCCDDD"
	fmt.Println(compress(a))
}
