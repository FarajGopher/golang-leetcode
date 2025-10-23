package main

import "fmt"

func fibonacci(n int) {
	a, b := 0, 1
	fmt.Println(a)
	if n > 1 {
		fmt.Println(b)
	}
	for i := 2; i <= n; i++ {
		a, b = b, a+b
		fmt.Println(b)
	}
}

func fibonacciRecusion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecusion(n-1) + fibonacciRecusion(n-2)
}

// Recursive Fibonacci sequence print
func fibonacciPrintRecursion(a, b, n int) {
	if n == 0 {
		return
	}
	fmt.Println(a)
	fibonacciPrintRecursion(b, a+b, n-1)
}


func main() {
	n := 10
	fmt.Println("Fibonacci of", n, "is:")
	fibonacci(n)
	println("Fibonacci of", n, "using recursion is", fibonacciRecusion(n))

	fmt.Println("\nFibonacci of", n, "terms (recursive print):")
	fibonacciPrintRecursion(0, 1, n)
}