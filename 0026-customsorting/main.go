package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

type byName []person

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {
	p1 := person{"Mark", 14}
	p2 := person{"Darwin", 25}
	p3 := person{"Jacoco", 7}
	p4 := person{"Pikachu", 18}
	p5 := person{"Ola", 2}

	people := []person{p1, p2, p3, p4, p5}

	fmt.Println("Original:", people)

	sort.Sort(byAge(people))

	fmt.Println("By Age Sorting: ", people)

	sort.Sort(byName(people))
	fmt.Println("By Name Sorting: ", people)
}
