/**
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
	first name
	last name
	favorite ice cream flavors

Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.


*/

package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	faveFlavors []string
}

func main() {
	p1 := person{
		firstName:   "Mary",
		lastName:    "WhiteCooper",
		faveFlavors: []string{"Chocolate", "Mintyliciousss"}}

	p2 := person{
		firstName:   "Ellen",
		lastName:    "Millstone",
		faveFlavors: []string{"Vanilla", "Just Ice", "Salted Egg"}}

	fmt.Println(p1)
	fmt.Println(p2)
}
