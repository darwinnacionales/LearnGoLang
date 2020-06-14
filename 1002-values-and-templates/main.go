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
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Darwin")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n\n------------------\n\n")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", "Mark")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n\n------------------\n\n")
}
