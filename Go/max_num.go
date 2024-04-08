package main

import (
	"fmt"
	"math"
)

func main() {

	//largest number
	var num1 float64 = 10.25
	var num2 float64 = 20.15
	var large float64 = 0
	large = math.Max(num1, num2)
	fmt.Printf("\nLargest number : %g\n", large)

	// Pow() function
	var num float64 = 10
	var p float64 = 4
	var result float64 = 0
	result = math.Pow(num, p)
	fmt.Printf("\nPower : %g\n", result)

	// math.Ceil()
	result = math.Ceil(num1)
	fmt.Printf("\nCeil value : %g\n", result)

	// math.Floor()
	result = math.Floor(num1)
	fmt.Printf("\nFloor value : %g\n", result)

	//Square root math.Sqrt()
	result = math.Sqrt(num1)
	fmt.Printf("\nSquare root = %g\n", result)
	fmt.Println()

}
