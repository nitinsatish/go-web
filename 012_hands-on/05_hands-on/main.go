package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

type item struct {
	Name  string
	Price float64
}
type menu struct {
	Timing string
	Item   item
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	mymenu := menu{
		Timing: "Breakfast",
		Item: item{
			Name:  "Oats",
			Price: 24.5,
		},
	}

	err := tpl.Execute(os.Stdout, mymenu)
	if err != nil {
		fmt.Println("Error !")
	}

}
