package main

import "fmt"

func main() {
	const num1 int = 1
	const num2 = 2

	const (
		num3 = 3
		num4 = 4
	)

	fmt.Println(num1, num2, num3, num4)
}
