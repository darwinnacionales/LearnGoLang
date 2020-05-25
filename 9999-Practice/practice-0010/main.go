/*
	Create a variable of type string using a raw string literal. Print it.
*/

package main

import "fmt"

func main() {
	str := `this is
	a raw
	"string literal"
	hehe`

	fmt.Println(str)
}
