package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

type person struct {
	Name        string
	YearOfBirth int
}

func (p person) GetCurrentAge() int {
	currentTime := time.Now()
	return currentTime.Year() - p.YearOfBirth
}

func (p person) TakesArgs(x int) int {
	return 5 + x
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p := person{
		"Darwin", 1995,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p)

	if err != nil {
		log.Fatalln(err)
	}
}
