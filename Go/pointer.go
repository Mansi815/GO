package main

import (
	"fmt"
	"math"
)

func main() {
	integer := 23
	floatingnum := 1234.573921109
	var val float64 = -19.25
	var ch byte = 'a'
	fmt.Printf("Absolute value of %.2f is %.2f\n", val, math.Abs(val))
	//bch := "B"
	fmt.Printf("%c\n", ch)
	fmt.Printf("ASCII VALUE : %d\n", ch)
	//fmt.Printf("%c", bch)
	fmt.Printf("%T\t%T\n", integer, &integer)
	fmt.Printf("5.2f : %5.2f\n", floatingnum)
	fmt.Printf("%*s\n", 15, "text")
	fmt.Printf("%040s", "text")
}
