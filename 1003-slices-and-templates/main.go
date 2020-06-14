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
	arr := []string{"Mark", "Darwin", "Chester", "Sam"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", arr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n\n------------------\n\n")
}
