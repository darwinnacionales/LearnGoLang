package main

import "fmt"

func main() {
	// Shorthand Declaration
	x := 5
	y := 3
	fmt.Println("Value of x is: ", x)

	z := x + y
	fmt.Println("Value of z is: ", z)

	name := "Darwin"
	fmt.Println("Value of name is: ", name)

	// Assigning a new value
	z = 5
	fmt.Println("New value of z is: ", z)

	name = "Pogi"
	fmt.Println("New value of name is: ", name)

	// Multiple variables of the same type
	multi1, multi2, multi3 := 100, 200, 300
	fmt.Println("Multi Same Types: ", multi1, multi2, multi3)

	// Multiple variables of different types
	multi4, multi5, multi6 := 400, "five hundred", 600
	fmt.Println("Multi Different Types: ", multi4, multi5, multi6)
}
