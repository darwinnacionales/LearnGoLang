package main

import "fmt"

func main() {
	// Shorthand
	str1 := "Hello"
	str2 := `
		This is
		a multiline

		string
	`

	// Explicit
	var str3 = "Hello"
	var str4 = `
		This is
		a multiline

		string
	`

	fmt.Println("Value of str1:", str1)
	fmt.Printf("Type of str1: %T\n\n", str1)

	fmt.Println("Value of str2:", str2)
	fmt.Printf("Type of str2: %T\n\n", str2)

	fmt.Println("Value of str3:", str3)
	fmt.Printf("Type of str3: %T\n\n", str3)

	fmt.Println("Value of str4:", str4)
	fmt.Printf("Type of str4: %T\n\n", str4)
}
