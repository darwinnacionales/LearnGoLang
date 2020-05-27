/*
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/
package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i <= 20 {
			fmt.Println(i, " less than or equal to 20")
		} else if i <= 50 {
			fmt.Println(i, " less than or equal to 50")
		} else {
			fmt.Println(i, " greater than 50")
		}
	}
}
