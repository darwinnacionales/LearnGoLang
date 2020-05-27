package main

import "fmt"

func main() {
	exampleOne()
	exampleTwo()
	exampleThree()
	exampleFour()
	exampleFive()
	exampleSix()
}

func exampleOne() {
	fmt.Println(" ----- Example One: -----")
	switch {
	case 1 == 1:
		fmt.Println("1 is equal to 1")
	case 2 == 2:
		fmt.Println("2 is equal to 2")
	case 3 == 3:
		fmt.Println("3 is equal to 3")
	case 2 == 3:
		fmt.Println("2 is not equal to 3")
	}
	fmt.Println()
}

func exampleTwo() {
	fmt.Println(" ----- Example Two: -----")
	switch {
	case 1 == 1:
		fmt.Println("1 is equal to 1")
		fallthrough
	case 2 == 2:
		fmt.Println("2 is equal to 2")
	case 3 == 3:
		fmt.Println("3 is equal to 3")
	case 2 == 3:
		fmt.Println("2 is not equal to 3")
	}
	fmt.Println()
}

func exampleThree() {
	fmt.Println(" ----- Example Three: -----")
	switch {
	case 1 == 1:
		fmt.Println("1 is equal to 1")
		fallthrough
	case 1 == 2:
		fmt.Println("1 is not equal to 2")
	case 3 == 3:
		fmt.Println("3 is equal to 3")
	case 2 == 3:
		fmt.Println("2 is not equal to 3")
	}
	fmt.Println()
}

func exampleFour() {
	fmt.Println(" ----- Example Four: -----")
	switch {
	case 1 == 2:
		fmt.Println("1 is not equal to 2")
	default:
		fmt.Println("Reached default")
	}
	fmt.Println()
}

func exampleFive() {
	fmt.Println(" ----- Example Five: -----")

	name := "Darwin"

	switch name {
	case "Darwin":
		fmt.Println("Name is Darwin")
	case "Test":
		fmt.Println("Name is Test")
	default:
		fmt.Println("Reached default")
	}
	fmt.Println()
}

func exampleSix() {
	fmt.Println(" ----- Example Six: -----")

	name := "Yehey"

	switch name {
	case "Darwin", "Mark", "Yehey":
		fmt.Println("Name is Darwin")
	case "Test":
		fmt.Println("Name is Test")
	default:
		fmt.Println("Reached default")
	}
	fmt.Println()
}
