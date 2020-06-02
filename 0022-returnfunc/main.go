package main

import "fmt"

func main() {
	funcRef := sample()
	y := funcRef()
	fmt.Println(y)

	addRef := add()
	x := addRef(5, 3)
	fmt.Println(x)
}

func sample() func() int {
	return func() int {
		return 28
	}
}

func add() func(a int, b int) int {
	return func(a int, b int) int {
		return a + b
	}
}
