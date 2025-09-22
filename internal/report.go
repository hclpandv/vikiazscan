package internal

import (
	"html/template"
	"os"
)

// GenerateHTMLReport renders multiple sections (one per KQL file)
func GenerateHTMLReport(outputFile string, data map[string]struct {
	Headers []string
	Rows    [][]string
}) error {
	tmpl, err := template.ParseFiles("templates/report.html")
	if err != nil {
		return err
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
