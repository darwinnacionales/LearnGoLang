/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/
package main

import "fmt"

func main() {
	favSport := "None"

	switch favSport {
	case "Basketball":
		fmt.Println("Oh yeahh")
	case "Baseball":
		fmt.Println("Yayyyy")
	default:
		fmt.Println("booooo")
	}
}
