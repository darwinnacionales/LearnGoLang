package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"firstName"`
	Last  string `json:"lastName"`
	Age   int    `json:"age"`
}

func main() {
	p1 := person{
		First: "Mark",
		Last:  "Darwin",
		Age:   12,
	}

	p2 := person{
		First: "Woohooo",
		Last:  "Hehe",
		Age:   33,
	}

	people := []person{p1, p2}

	fmt.Println("Original:", people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON Format:", string(bs))

	var convertedPeople []person

	err = json.Unmarshal(bs, &convertedPeople)

	if err == nil {
		fmt.Println("Back to Original: ", convertedPeople)
	}
}
