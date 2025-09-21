package cmd

import (
	"fmt"
	"log"

	"github.com/hclpandv/vikiazscan/internal"
	"github.com/spf13/cobra"
)

var ScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for orphaned resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		resources, err := internal.ExecuteKQLQuery("queries/orphan_disks.kql")
		if err != nil {
			log.Fatalf("Error scanning: %v", err)
		}

		err = internal.GenerateHTMLReport("vikiazscan-report.html", resources)
		if err != nil {
			log.Fatalf("Error generating report: %v", err)
		}

		fmt.Println("HTML report generated: vikiazscan-report.html")
		return nil
	},
}
