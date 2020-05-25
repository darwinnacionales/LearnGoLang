/*
	Building on the code from the previous example
	at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
	The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
	eg:

	in func main
		this should already be done
		print out the value of the variable “x”
		print out the type of the variable “x”
		assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
		print out the value of the variable “x”

		now do this
		now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
		then use the “=” operator to ASSIGN that value to “y”
		print out the value stored in “y”
		print out the type of “y”

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

	// now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
	//	then use the “=” operator to ASSIGN that value to “y”
	var y int = int(x)

	// print out the value of the variable “y”
	fmt.Println(y)

	// print out the type of the variable “y”
	fmt.Printf("%T\n", y)
}
