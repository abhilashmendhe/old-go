package main

import (
	"html/template"
	"log"
	"os"
)

type Invoice struct {
	Name    string
	Paid    bool
	Charges []float64
	Total   float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	html, err := template.ParseFiles("bill.html")

	bill := Invoice{
		Name:    "Mary Gibbs",
		Paid:    true,
		Charges: []float64{23.1, 1.13, 42.28},
		Total:   74.23,
	}
	err = html.Execute(os.Stdout, bill)
	check(err)
}
