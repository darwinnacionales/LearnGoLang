/*
	Create TYPED and UNTYPED constants. Print the values of the constants.
*/

package main

import "fmt"

func main() {
	const (
		a int = 5
		b     = 3
	)

	fmt.Println(a, b)
}
