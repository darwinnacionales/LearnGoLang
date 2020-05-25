/*
	Using the following operators, write expressions and assign their values to variables:
			==
			<=
			>=
			!=
			<
			>
	Now print each of the variables.

*/

package main

import "fmt"

func main() {
	a := "hello" == "hello"
	b := 3 <= 6
	c := 3 >= 6
	d := 3 != 6
	e := 3 < 6
	f := 3 > 6

	fmt.Println(a, b, c, d, e, f)
}
