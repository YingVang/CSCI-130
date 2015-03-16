// source:  http://golang.org/pkg/text/template/#example_Template

package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Honorific}}{{.LastName}},
{{if .Attended}}
It was a pleasure to see you at the fundraiser.{{else}}
It is a shame you couldn't attend the fundraiser.{{end}}
{{if .Donated}}
Thank you for donating.{{end}}
{{range .UpcommingEvents}}
    Remember the upcoming events: {{.}}
{{end}}
Best wishes,
Jerry
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Honorific, LastName string
		Attended, Donated   bool
		UpcommingEvents		[]string
	}
	
	var upcommingEvents = []string{"Face Painting",
		"Water Ballon Fight",
		"Tug of War"}	
		
	var recipients = []Recipient{
		{"Mr", " Jones", true, true, upcommingEvents},
		{"Mrs", " Mildred", false, true, upcommingEvents},
		{"Mr", " Rodney", true, false, upcommingEvents},
		{"Ms", " Wilson", false, false, upcommingEvents},
	}
		
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))


	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
	
	fmt.Println("")

}