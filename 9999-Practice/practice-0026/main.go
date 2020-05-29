/**
Create a map with a key of TYPE string which is a person’s “last_first” name,
and a value of TYPE []string which stores their favorite things. Store three records in your map.
Print out all of the values, along with their index position in the slice.

	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
	`no_dr`, `Being evil`, `Ice cream`, `Sunsets`



*/

package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	m["Wahaha"] = []string{"to laugh", "be laughed at", "chuckle"}

	for k, arr := range m {
		fmt.Println(k)

		for i, v := range arr {
			fmt.Println(i, v)
		}
		fmt.Println()
	}

	delete(m, "Wahaha")

	for k, arr := range m {
		fmt.Println(k)

		for i, v := range arr {
			fmt.Println(i, v)
		}
		fmt.Println()
	}
}
