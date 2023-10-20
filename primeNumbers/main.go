package main

import (
	"errors"
	"fmt"
	"strconv"
)

func getPrimeInput() (int, error) {
	var primeNumberInput string

	fmt.Print("Type the prime number target: ")
	_, err := fmt.Scan(&primeNumberInput)
	if err != nil {
		return 0, errors.New("please, type a valid number")
	}

	primeNum, err := strconv.Atoi(primeNumberInput)
	if err != nil {
		return 0, errors.New("please, type a valid number")
	}

	if primeNum < 2 {
		return 0, errors.New("please, type a valid number greater than 2")
	}

	return primeNum, nil
}

func sieveOfEratosthenesRecursive(N int, current int) []int {
	if current > N {
		return []int{}
	}

	primes := []int{}
	if isPrime(current, 2) {
		primes = append(primes, current)
	}

	restPrimes := sieveOfEratosthenesRecursive(N, current+1)
	return append(primes, restPrimes...)
}

func isPrime(num int, divisor int) bool {
	if num <= 2 {
		return num == 2
	}
	if num%divisor == 0 {
		return false
	}
	if divisor*divisor > num {
		return true
	}
	return isPrime(num, divisor+1)
}

func sieveOfEratosthenesLinear(N int) []int {

	numbers := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		numbers[i] = true
	}

	for i := 2; i*i <= N; i++ {
		if numbers[i] {
			for j := i * i; j <= N; j += i {
				numbers[j] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i <= N; i++ {
		if numbers[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	primeNum, err := getPrimeInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n[recursivePrime] | prime(%v): Value: %v", primeNum, sieveOfEratosthenesRecursive(primeNum, 2))
	fmt.Printf("\n[linearPrime]    | prime(%v): Value: %v", primeNum, sieveOfEratosthenesLinear(primeNum))
}
