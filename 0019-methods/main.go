package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) displayAgeAfterNYears(n int) {
	fmt.Printf("%v is %d years old after %d years.", p.name, p.age+n, n)
}

func main() {
	p1 := person{name: "Omar", age: 17}
	p1.displayAgeAfterNYears(23)
}
