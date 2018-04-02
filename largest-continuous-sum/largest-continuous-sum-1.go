package main

import "fmt"

var a []int

func LargestContinuousSum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	currentSum := a[0]
	maxSum := a[0]

	for i := 1; i < len(a); i++ {
		currentSum += i
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	fmt.Scan(&a)
	fmt.Println(LargestContinuousSum(a))
}
