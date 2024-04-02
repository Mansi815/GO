package main

import "fmt"

const Pi = 3.14

func main() {
	const LENGTH int = 10
	const WIDTH int = 10
	const Const int = 1
	var area1 int
	var radius float64

	var a uint16
	a = 300
	area := 0.0
	perimeter := 0.0

	fmt.Printf("Enter the radius : \n")
	fmt.Scanf("%g", &radius)

	// Area of circle
	area = Pi * radius * radius
	fmt.Printf("Area of circle %g\n", area)

	// Perimeter of circle
	perimeter = 2 * Pi * radius
	fmt.Printf("Perimeter of circle %g\n", perimeter)

	// Area of Rectangle
	fmt.Printf("a= %d\n", a)
	area1 = LENGTH * WIDTH
	fmt.Printf("Area of Rectangle : %d\n", area1)

	//Constant
	fmt.Println(Const)

	// Conversion FRom F to C
	
}
