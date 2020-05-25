package main

import "fmt"

func main() {
	// Shorthand
	num1 := 1
	num2 := 2.56

	// Explicit
	var num3 int = 3
	var num4 float64 = 4.6446

	fmt.Println("num1 value:", num1)
	fmt.Printf("num1 type: %T\n", num1)

	fmt.Println("num2 value:", num2)
	fmt.Printf("num2 type: %T\n", num2)

	fmt.Println("num3 value:", num3)
	fmt.Printf("num3 type: %T\n", num3)

	fmt.Println("num4 value:", num4)
	fmt.Printf("num4 type: %T\n", num4)
}
