/*
	https://golang.org/ref/spec#Types
	For this exercise
	Create your own type. Have the underlying type be an int.
	create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword

	in func main
		print out the value of the variable “x”
		print out the type of the variable “x”
		assign 42 to the VARIABLE “x” using the “=” OPERATOR
		print out the value of the variable “x”
*/

package main

import "fmt"

// Create your own type. Have the underlying type be an int.
type mytype int

// create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
var x mytype

func main() {
	// print out the value of the variable “x”
	fmt.Println(x)

	// print out the type of the variable “x”
	fmt.Printf("%T\n", x)

	// assign 42 to the VARIABLE “x” using the “=” OPERATOR
	x = 42

	// print out the value of the variable “x”
	fmt.Println(x)
}
