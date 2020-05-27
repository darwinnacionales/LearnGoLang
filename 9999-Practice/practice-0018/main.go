/*
Create a program that uses a switch statement with no switch expression specified.
*/
package main

import "fmt"

func main() {
	switch {
	case 1 == 1:
		fmt.Println("1 is equal to 1")
	default:
		fmt.Println("This is default")
	}
}
