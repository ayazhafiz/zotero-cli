package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Zotero CLI configuration",
	Long:  `Supplies utilities for managing the Zotero CLI configuration.`,
}

func init() {
	rootCmd.AddCommand(configCmd)
}
