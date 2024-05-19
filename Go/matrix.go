package main

import "fmt"

func main() {
	var matrix [3][3]int
	fmt.Printf("Enter 9 matrix elements:\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Element [%d][%d]: ", i, j)
			fmt.Scanf("%d", &matrix[i][j])
		}
	}
	fmt.Println("Matrix:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
