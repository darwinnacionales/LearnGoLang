/*
	Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
		42
		“James Bond”
		true

	Now print the values stored in those variables using
		a single print statement
		multiple print statements
*/

package main

import "fmt"

func main() {
	// ASSIGN these VALUES to VARIABLES
	x, y, z := 42, "James Bond", true

	// print - a single print statement
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// print - multiple print statements
	fmt.Println(x, y, z)
}
