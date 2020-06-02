package main

import "fmt"

func main() {
	displayName()
	displaySum(4, 5)

	x := multiply(10, 2)
	fmt.Println(x)

	name, age := getNameAge()
	fmt.Println(name, age)
}

func displayName() {
	fmt.Println("Darwin")
}

func displaySum(x int, y int) {
	fmt.Println(x + y)
}

func multiply(x int, y int) int {
	return x * y
}

func getNameAge() (string, int) {
	name := "Brad"
	age := 54
	return name, age
}
