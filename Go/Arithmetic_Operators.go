//Arithmetic Operations

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var n, n1, n2, sum, diff, mul, div, mod int

	n1 = 30
	n2 = 20

	sum = n1 + n2
	diff = n1 - n2
	mul = n1 * n2
	div = n1 / n2
	mod = n1 % n2
	// All the Arithmetic Operatores
	fmt.Printf("Sum of %d and %d is %d\n", n1, n2, sum)
	fmt.Printf("Difference of %d and %d is %d\n", n1, n2, diff)
	fmt.Printf("Muliplication of %d and %d is %d\n", n1, n2, mul)
	fmt.Printf("Division of %d and %d is %d\n", n1, n2, div)
	fmt.Printf("Modulus of %d and %d is %d\n", n1, n2, mod)
	n1++
	fmt.Println("Increment of n1 is", n1)
	n1--
	fmt.Println("Decrement of n1 is", n1)
	n2 += 5
	fmt.Printf("After adding 5 to 20 : %d \n ", n2)
	// & and | or
	a := 5
	b := 3
	var m float32 = 10.3
	fmt.Printf("Size of a = %d\n", unsafe.Sizeof(a))
	fmt.Printf("Size of m = %d\n", unsafe.Sizeof(m))

	c := a & b
	fmt.Printf("& Operator of %d and %d is %d\n", a, b, c)
	c = a | b
	fmt.Printf("| Operator of %d and %d is %d\n", a, b, c)
	c = a ^ b
	fmt.Printf("XOR Operator of %d and %d is %d\n", a, b, c)
	s := 5
	fmt.Println(s > 5)
	fmt.Println(s == 5)
	fmt.Println(s < 5)

	s <<= 2
	fmt.Printf("Left shift of s(5) :%d", s)
	fmt.Printf("%d\n", n)
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &n)

	mod = n % 2
	fmt.Printf("Number is even or odd?(1=odd and 0=even) \n%d", mod)
}
