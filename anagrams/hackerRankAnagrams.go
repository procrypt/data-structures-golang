package main

import (
	"fmt"
	"math"
)

var (
	 a string
	 b string
	 )

func abs(num int) int {
	return int(math.Abs(float64(num)))
}

func main() {

	fmt.Scan(&a)
	fmt.Scan(&b)

	deletion := 0

	// To make sure length of each string is equal
	alphabet := make([]int, 26)

	for _, char := range a {
		alphabet[int(char)-int('a')]++
	}

	for _, char := range b {
		alphabet[int(char)-int('a')]--
	}

	for i := 0; i < 26; i++ {
		deletion += abs(alphabet[i])
	}


}