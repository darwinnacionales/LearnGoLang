package main

import "fmt"

func main() {
	// Key - type string
	// Value - type int
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Println(m)

	// Accessing data
	fmt.Println("Result of looking for one", m["one"])
	fmt.Println("Result of looking for two", m["two"])
	fmt.Println("Result of looking for three", m["three"])

	// Checking if item exists
	val, isExisting := m["one"]
	fmt.Println(val, isExisting)

	if val, isExisting := m["two"]; isExisting {
		fmt.Printf("%d exists\n", val)
	}

	// Add item in map
	m["three"] = 3
	fmt.Println("Result of looking for three after adding it", m["three"])

	// Loop through map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Delete item from map
	delete(m, "three")
	fmt.Println(m)
}
