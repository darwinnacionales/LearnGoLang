package main

import "fmt"

func main() {
	defer func1()
	func2()
	fmt.Println("End of main()")
}

func func1() {
	fmt.Println("This is from func1()")
}

func func2() {
	fmt.Println("This is from func2()")
}
