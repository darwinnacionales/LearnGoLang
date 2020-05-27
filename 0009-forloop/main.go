package main

import "fmt"

func main() {
	fmt.Println("-------- Basic For Loop --------------")

	// For Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------- Break and Continue --------------")

	// Break and continue
	for i := 0; i < 10; i++ {
		if i <= 3 {
			continue
		}

		if i >= 7 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("-------- While Loop in GoLang --------------")

	// while equivalent
	x := 2
	for x < 10 {
		fmt.Println(x)
		x *= 2
	}

	fmt.Println("---------- Infinite Loop ------------")

	// Inifinite loop
	y := 1
	for {
		if y > 10 {
			break
		}
		fmt.Println(y)
		y *= 2
	}
}
