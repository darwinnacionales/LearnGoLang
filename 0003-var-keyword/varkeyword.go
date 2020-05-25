package main

import "fmt"

// Global variable
var z = 40
var w int = 30

func main() {
	// Local Variable
	var x = 15
	var y int = 20

	fmt.Println(w, x, y, z)

	// Reassigning value to z
	z = 90

	// Local variable is given preference over global variable
	w := 70

	fmt.Println(w, x, y, z)

	displayGlobalVariables()
}

func displayGlobalVariables() {
	fmt.Println(w, z)
}
