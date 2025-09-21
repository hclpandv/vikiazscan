package internal

import (
	"html/template"
	"os"
)

func GenerateHTMLReport(filename string, resources []OrphanedResource) error {
	tmpl, err := template.ParseFiles("templates/report.html")
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, resources)
}
