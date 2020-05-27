/*
Create a program that shows the “if statement” in action.
*/
package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i >= 20 {
			break
		}

		fmt.Println(i)
	}
}
