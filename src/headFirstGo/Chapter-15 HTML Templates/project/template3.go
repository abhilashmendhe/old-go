package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {

	tmpl, err := template.New("testing").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func main() {

	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"

	executeTemplate(templateText, []string{"do", "re", "mi"})
	executeTemplate("Dot is:{{.}}\n", "80 harzar")
	executeTemplate("Dot is:{{.}}\n", "bate mat karo bade bade")

	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", false)

	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 0.99, 28})
	executeTemplate(templateText, nil)
	executeTemplate(templateText, []float64{})

	templateText = "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})

	templateText = "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	sub := Subscriber{Name: "Abhilash", Rate: 2.1, Active: true}
	executeTemplate(templateText, sub)
	sub = Subscriber{Name: "Abhi", Rate: 10.2, Active: false}
	executeTemplate(templateText, sub)
}
