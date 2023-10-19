# **Fibonacci Calculator in GoLang**

This is a simple GoLang program that calculates Fibonacci numbers using two different approaches.

## **How to Use**

1. Ensure you have GoLang installed on your system.
2. Run the following command to execute the program:
   ```bash
   go run main.go
3. Enter the number of Fibonacci numbers you want to calculate when prompted.

## **Functions**

### `recursiveFib`

This function calculates Fibonacci numbers recursively. However, it's not efficient for large numbers due to its recursive nature. So, there's a special case when "n" exceeds a certain limit, where, due to the inefficiency of the recursive algorithm, the program will return a -1.

### `linearFib`

This function uses memoization to optimize the Fibonacci calculation. It stores previously calculated Fibonacci numbers in a map to avoid redundant calculations and significantly speed up the calculation process.

### `getFibonacciInput`

This function is designed to obtain user input for the number of Fibonacci sequences they want to generate. It ensures that the input is a positive integer and returns the entered number as an integer or an error if the input is invalid.

## Be careful with [Stack Overflow](https://www.techtarget.com/whatis/definition/stack-overflow#:~:text=A%20stack%20overflow%20is%20a,been%20allocated%20to%20that%20stack.)!
