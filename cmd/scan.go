package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/hclpandv/vikiazscan/internal"
	"github.com/spf13/cobra"
)

var ScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for orphaned resources",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Discover all .kql files in queries folder
		queries := internal.GetKQLFiles("queries")

		// Map for tabs: folder/category -> list of tables
		tabs := make(map[string][]internal.TableData)

		for _, kqlFile := range queries {
			headers, rows, err := internal.ExecuteKQLQuery(kqlFile)
			if err != nil {
				log.Printf("Error scanning %s: %v", kqlFile, err)
				continue
			}

			// derive category from folder
			rel, err := filepath.Rel("queries", kqlFile)
			if err != nil {
				log.Printf("Error getting relative path for %s: %v", kqlFile, err)
				continue
			}
			folder := filepath.Dir(rel)
			if folder == "." {
				folder = "uncategorized"
			}

			// derive table name from file name
			base := filepath.Base(kqlFile)
			name := strings.TrimSuffix(base, ".kql")
			name = strings.ReplaceAll(name, "_", " ")

			// append table to the appropriate tab
			table := internal.TableData{
				Name:    name,
				Headers: headers,
				Rows:    rows,
			}
			tabs[folder] = append(tabs[folder], table)
		}

		// Generate HTML report
		err := internal.GenerateHTMLReport("vikiazscan-report.html", tabs)
		if err != nil {
			log.Fatalf("Error generating report: %v", err)
		}

		fmt.Println("HTML report generated: vikiazscan-report.html")
		return nil
	},
}
