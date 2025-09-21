package main

import (
	"os"

	"github.com/hclpandv/vikiazscan/cmd"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{Use: "vikiazscan"}
	root.AddCommand(cmd.ScanCmd)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
