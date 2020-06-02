package main

import "fmt"

func main() {
	func1 := func() {
		fmt.Println("Inside func 1")
	}

	func1()

	func2 := func(x int) {
		fmt.Println("I received:", x)
	}

	func2(75)
}
