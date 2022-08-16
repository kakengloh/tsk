package cmd

import "github.com/spf13/cobra"

var edCmd = &cobra.Command{
	Use:   "ed",
	Short: "Edit an existing task",
	Run:   func(cmd *cobra.Command, args []string) {},
}
