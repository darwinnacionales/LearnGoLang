package main

import "fmt"

func main() {
	func() {
		fmt.Println("Inside anon func 1")
	}()

	func(x int) {
		fmt.Println("I received:", x)
	}(10)
}
