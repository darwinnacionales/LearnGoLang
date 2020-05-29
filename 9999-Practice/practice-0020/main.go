/**
Using a COMPOSITE LITERAL:
	create a SLICE of TYPE int
	assign 10 VALUES
	Range over the slice and print the values out.
	Using format printing
		print out the TYPE of the slice
*/

package main

import "fmt"

func main() {
	x := [5]int{}

	x[0] = 0
	x[1] = 1
	x[2] = 2
	x[3] = 3
	x[4] = 4

	for _, v := range x {
		fmt.Printf("%d\t", v)
	}

	fmt.Printf("\n%T\n", x)
}
