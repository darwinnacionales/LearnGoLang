/*
	Write a program that prints a number in decimal, binary, and hex
*/

package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t\t%b\t\t%#x\n", i, i, i)
	}
}
