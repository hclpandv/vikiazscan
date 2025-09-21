package internal

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

func GenerateHTMLReport(filename string, resources []OrphanedResource) error {
	// Load the template from the file
	t, err := template.ParseFiles("templates/report.html")
	if err != nil {
		return fmt.Errorf("failed to parse HTML template: %v", err)
	}

	// Data for the template
	data := struct {
		Date      string
		Resources []OrphanedResource
	}{
		Date:      time.Now().Format("2006-01-02 15:04:05"),
		Resources: resources,
	}

	// Create the output file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create HTML report: %v", err)
	}
	defer file.Close()

	// Execute the template and write to the file
	err = t.Execute(file, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}
