package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/tpl1.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type product struct {
	ProductName string
	Price       int
}

func main() {
	prod := product{
		ProductName: "iPhone 12",
		Price:       2500,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", prod)
	if err != nil {
		log.Fatalln(err)
	}
}
