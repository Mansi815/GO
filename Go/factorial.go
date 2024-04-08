package main

import "fmt"

var n int

func factorial(n int) int {

	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	fmt.Printf("Enter a number = ")
	fmt.Scanf("%d\n", &n)
	result := factorial(n)
	fmt.Printf("Factorial of %d = %d", n, result)
}
