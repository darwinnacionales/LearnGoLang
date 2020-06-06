package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 9, 1, 0, 4, 2, 8, 3, 6}
	xs := []string{"cranberry", "pineapple", "apple", "blueberry", "orange", "lemon", "grapes"}

	fmt.Println("Unsorted")
	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("Sorted")
	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
