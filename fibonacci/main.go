package main

import (
	"fmt"
)

var memo = make(map[int]int)

func recursiveFib(n int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else if n >= 50 {
		return -1
	}
	return recursiveFib(n-1) + recursiveFib(n-2)
}

func linearFib(n int) int {
	if memo[n] != 0 {
		return memo[n]
	} else if n <= 2 {
		return 1
	}

	memo[n] = (linearFib(n-1) + linearFib(n-2))
	return memo[n]
}

func main() {
	var fibNumber int
	fmt.Print("Type the number of Fibonacci you want: ")
	fmt.Scan(&fibNumber)

	fmt.Printf("\n[recursiveFib] | fib(%v): Value: %v", fibNumber, recursiveFib(fibNumber))
	fmt.Printf("\n[linearFib]    | fib(%v): Value: %v", fibNumber, linearFib(fibNumber))
}
