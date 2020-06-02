package main

import "fmt"

func main() {
	displayNums(1, 2, 3, 4, 5)

	x := []int{10, 11, 12, 13, 14, 15}
	displayNums(x...)
}

func displayNums(x ...int) {
	fmt.Println(x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("The total is", sum)
}
