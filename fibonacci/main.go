package main

import (
	"fmt"
)

func recursiveFib(n int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else if n > 40 {
		return -1
	}
	return recursiveFib(n-1) + recursiveFib(n-2)
}

func main() {
	var fibNumber int
	fmt.Print("Type the number of Fibonacci you want: ")
	fmt.Scan(&fibNumber)

	fmt.Printf("\n[recursiveFib] | fib(%v): Value: %v", fibNumber, recursiveFib(fibNumber))
}
