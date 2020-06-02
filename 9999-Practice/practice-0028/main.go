/**
Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
Access each value in the map.
Print out the values, ranging over the slice.
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

	m := map[string]person{
		"Mary": p1,
	}

	m["Ellen"] = p2

	for _, person := range m {
		fmt.Println(person.firstName, person.faveFlavors)
	}
}
