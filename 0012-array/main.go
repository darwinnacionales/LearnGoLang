package main

import "fmt"

func main() {
	var x [6]int
	y := [5]int{1, 2, 3, 4, 5}

	// Display contents of array
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// Update by index
	x[3] = 42
	fmt.Println("Updated x:", x)

	// Display length
	fmt.Println("Length of x:", len(x))
}
