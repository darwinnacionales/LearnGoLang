/*
	Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
*/

package main

import "fmt"

const (
	year0 = 2020 + iota
	year1
	year2
	year3
)

func main() {
	fmt.Println(year0, year1, year2, year3)
}
