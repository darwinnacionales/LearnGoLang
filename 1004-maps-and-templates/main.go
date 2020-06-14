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

func main() {
	m := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
		"4": "four",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", m)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n\n------------------\n\n")
}
