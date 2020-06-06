package main

import "fmt"

func main() {
	a := 42
	b := &a
	fmt.Println(a, "with address", b)

	fmt.Println("Original value is", *b)

	*b = 900
	fmt.Println("a is now updated to", a)

	fmt.Println("--------------------------")

	x := 0
	fmt.Println("Value of x is:", x)
	updateValue(&x)
	fmt.Println("New value of x is:", x)
}

func updateValue(y *int) {
	*y = 35
}
