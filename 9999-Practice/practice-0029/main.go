package main

/*
Create a value and assign it to a variable.
Print the address of that value.
*/

import "fmt"

func main() {
	a := 80
	fmt.Println("Address of a is:", &a)
}
