# **Prime Number Calculator in GoLang**

This is a simple GoLang program that calculates all the prime numbers using two different approaches, one recursive and one linear.

## **How to Use**

1. Ensure you have GoLang installed on your system.
2. Run the following command to execute the program:
   ```bash
   go run main.go
3. Enter the number of Fibonacci numbers you want to calculate when prompted.

## **Functions**

### `getPrimeInput`

This function retrieves a target prime number input from the user. It checks if the input is a valid integer greater than 2 and returns the entered number or an error message.

### `sieveOfEratosthenesRecursive(N, current)`

This recursive function implements the Sieve of Eratosthenes algorithm to find prime numbers up to a given target number, 'N'. It starts from 'current' and returns a slice containing all prime numbers found.

### `isPrime(num, divisor)`

A helper function used by the recursive sieve function to determine whether a given number 'num' is prime by recursively checking its divisibility with increasing divisors.

### `sieveOfEratosthenesLinear(N)`

This function implements the Sieve of Eratosthenes algorithm using a linear approach. It efficiently finds all prime numbers up to 'N' by marking and filtering non-prime numbers in an array. It returns a slice of prime numbers.

## Be careful with [Stack Overflow](https://www.techtarget.com/whatis/definition/stack-overflow#:~:text=A%20stack%20overflow%20is%20a,been%20allocated%20to%20that%20stack.)!
