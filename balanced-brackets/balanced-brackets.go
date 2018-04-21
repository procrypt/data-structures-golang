package main

import (
	"fmt"
	"os"
	"bufio"
)

var a int

type Stack struct {
	container []interface{}
}

func(s *Stack) Push(item interface{}) {
	s.container = append(s.container, item)
}

func(s *Stack) Pop() interface{} {
	if len(s.container) > 0 {
		s.container = s.container[:len(s.container)-1]
		return s.container
	} else {
		return fmt.Sprint("Empty Stack!!")
	}
}

func(s *Stack) Top() interface{} {
	if len(s.container) > 0 {
		return s.container[len(s.container)-1]
	} else {
		return fmt.Sprint("Empty Stack!!")
	}
}

func(s *Stack) Size() int {
	return len(s.container)
}

func stack() *Stack{
	return &Stack{}
}

func BalanceCheck(expression string) string {
	a := stack()
	for i := range expression {
		if string(expression[i]) == "{" ||
			string(expression[i]) == "[" ||
			string(expression[i]) == "("  {
			a.Push(string(expression[i]))
		} else {
			if a.Size() != 0 {
				if  string(expression[i]) == ")" && a.Top() == "("  ||
					string(expression[i]) == "]" && a.Top() == "["  ||
					string(expression[i]) == "}" && a.Top() == "{" {
					a.Pop()
				}
			}
		}
	}
	if a.Size() != 0 {
		return "NO"
	} else {
		return "YES"
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	fmt.Scan(&a)
	for i:=0; i<a; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println(BalanceCheck(text))
		}
	}
}