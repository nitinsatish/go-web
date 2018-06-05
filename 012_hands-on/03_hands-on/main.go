package main

import (
	"fmt"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	hlist := hotels{
		hotel{
			Name:    "Hotel California",
			Address: "#201, Near MG road",
			City:    "Bangalore",
			Zip:     560034,
			Region:  "Central",
		},
		hotel{
			"Hotel Taj", "#112 , Trinity circle", "Bengaluru", 450034, "South",
		},
	}
	err := tpl.Execute(os.Stdout, hlist)
	if err != nil {
		fmt.Println("Error!")
	}

}
