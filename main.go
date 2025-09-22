package main

import (
	"github.com/hclpandv/vikiazscan/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "vikiazscan",
		Short: "vikiazscan CLI for Azure resource assessment",
	}

	rootCmd.AddCommand(cmd.ScanCmd)
	rootCmd.Execute()
}
