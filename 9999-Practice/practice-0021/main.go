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
	x := [10]int{}

	for i := range x {
		x[i] = i * 10
	}

	for _, v := range x {
		fmt.Printf("%d\t", v)
	}

	fmt.Println()
}
