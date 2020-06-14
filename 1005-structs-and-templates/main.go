package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type product struct {
	ProductName string
	Price       int
}

func main() {
	p := product{
		ProductName: "Galaxy Note 20",
		Price:       1599,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n\n------------------\n\n")
}
