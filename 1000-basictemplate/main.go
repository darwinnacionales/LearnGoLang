package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n---------------------------\n\n")

	templates, err := template.ParseFiles("tpl.gohtml", "one.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = templates.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n---------------------------\n\n")
}
