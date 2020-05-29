package main

import "fmt"

func main() {
	x := make([]int, 2, 3)
	fmt.Printf("Contents: %v | Length: %d | Capacity: %d\n", x, len(x), cap(x))

	x = append(x, 1)
	fmt.Printf("Contents: %v | Length: %d | Capacity: %d\n", x, len(x), cap(x))

	x = append(x, 2)
	fmt.Printf("Contents: %v | Length: %d | Capacity: %d\n", x, len(x), cap(x))

	x = append(x, 3, 4)
	fmt.Printf("Contents: %v | Length: %d | Capacity: %d\n", x, len(x), cap(x))

	x = append(x, 5)
	fmt.Printf("Contents: %v | Length: %d | Capacity: %d\n", x, len(x), cap(x))
}

func basics() {
	// original array
	origArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("origArray: ", origArray)

	sliceSample1 := origArray[3:7]
	fmt.Println("sliceSample1: ", sliceSample1)

	// Slice Defaults
	sliceSample2 := origArray[:3]
	fmt.Println("sliceSample2: ", sliceSample2)

	sliceSample3 := origArray[8:]
	fmt.Println("sliceSample3: ", sliceSample3)

	// Updates
	sliceSample1[0] = 100
	fmt.Println("Updated sliceSample1: ", sliceSample1)
	fmt.Println("Updated origArray: ", origArray)

	// Length and Capacity
	fmt.Println("Length of sliceSample1:", len(sliceSample1))
	fmt.Println("Capacity of sliceSample1:", cap(sliceSample1))
}

func loops() {
	// original array
	origArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range origArray {
		fmt.Println(i, origArray[i])
	}

	fmt.Println("---------------------")

	for i, v := range origArray {
		fmt.Println(i, v)
	}
}

func appendAndDelete() {
	// original array
	x := []int{1, 2, 3}
	y := []int{7, 8, 9}
	fmt.Println("x: ", x)

	// Appending values
	x = append(x, 4, 5, 6)
	fmt.Println("x: ", x)

	// Appending a slice
	x = append(x, y...)
	fmt.Println("x: ", x)

	// Deletion
	x = append(x[:2], x[4:]...)
	fmt.Println("x: ", x)
}
