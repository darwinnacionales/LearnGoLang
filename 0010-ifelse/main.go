package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 0 || i == 1 {
			fmt.Print("zero or one")
		} else if i == 2 {
			fmt.Print("two")
		} else if i == 3 {
			fmt.Print("three")
		} else if i <= 6 {
			fmt.Print(i)
		} else {
			fmt.Print("More than 6")
		}

		fmt.Print(" ")

		if i%2 == 0 {
			fmt.Println(" even")
		} else {
			fmt.Println(" odd")
		}
	}
}
