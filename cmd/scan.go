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

		results := make(map[string]struct {
			Headers []string
			Rows    [][]string
		})

		for _, kqlFile := range queries {
			headers, rows, err := internal.ExecuteKQLQuery(kqlFile)
			if err != nil {
				log.Printf("Error scanning %s: %v", kqlFile, err)
				continue
			}

			// derive section name from file name
			base := filepath.Base(kqlFile)
			name := strings.TrimSuffix(base, ".kql")
			name = strings.ReplaceAll(name, "_", " ")

			results[name] = struct {
				Headers []string
				Rows    [][]string
			}{headers, rows}
		}

		// Generate HTML report
		err := internal.GenerateHTMLReport("vikiazscan-report.html", results)
		if err != nil {
			log.Fatalf("Error generating report: %v", err)
		}

		fmt.Println("HTML report generated: vikiazscan-report.html")
		return nil
	},
}
