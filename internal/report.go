package internal

import (
	"html/template"
	"os"
	"strings"
)

// TableData represents a single query table
type TableData struct {
	Name    string
	Headers []string
	Rows    [][]string
}

// GenerateHTMLReport renders multiple tabs (one per folder/category)
func GenerateHTMLReport(outputFile string, tabs map[string][]TableData) error {
	funcMap := template.FuncMap{
		"title": func(s string) string {
			// Convert folder names like "orphaned-resources" -> "Orphaned Resources"
			return strings.Title(strings.ReplaceAll(s, "-", " "))
		},
	}

	tmpl, err := template.New("report.html").Funcs(funcMap).ParseFiles("templates/report.html")
	if err != nil {
		return err
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, tabs)
}
