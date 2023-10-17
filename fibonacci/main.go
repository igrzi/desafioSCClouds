package main

import (
	"errors"
	"fmt"
	"strconv"
)

var memo = make(map[int]int)

func getFibonacciInput() (int, error) {
	var fibNumberInput string

	fmt.Print("Type the number of Fibonacci you want: ")
	fmt.Scan(&fibNumberInput)

	fibNumber, err := strconv.Atoi(fibNumberInput)
	if err != nil {
		return 0, errors.New("please enter a valid integer")
	}

	if fibNumber < 0 {
		return 0, errors.New("please enter a positive integer")
	}

	return fibNumber, nil
}

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
	fibNumber, err := getFibonacciInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\n[recursiveFib] | fib(%v): Value: %v", fibNumber, recursiveFib(fibNumber))
	fmt.Printf("\n[linearFib]    | fib(%v): Value: %v", fibNumber, linearFib(fibNumber))
}
