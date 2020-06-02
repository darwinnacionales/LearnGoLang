package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type employee struct {
	person
	role string
}

func main() {
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{firstName: "Mark", lastName: "Darwin"}

	fmt.Println(p1.firstName, p1.lastName, p1.age)
}

func embeddedSample() {
	p1 := person{firstName: "Mark", lastName: "Darwin"}
	emp1 := employee{p1, "CEO"}

	emp2 := employee{
		person: person{"Conan", "O'brien", 43},
		role:   "CTO"}

	fmt.Println(emp1)
	fmt.Println(emp2)

	fmt.Println(emp1.person.firstName, emp1.role)
}

func basic() {
	p1 := person{firstName: "Mark", lastName: "Darwin"}
	fmt.Println(p1.firstName, p1.lastName, p1.age)

	p2 := person{"Conan", "O'brien", 43}
	fmt.Println(p2.firstName, p2.lastName, p2.age)

	p3 := new(person)
	p3.firstName = "Tony"
	p3.lastName = "Stark"
	fmt.Println(p3.firstName, p3.lastName, p3.age)
}
