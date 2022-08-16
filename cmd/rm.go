package cmd

import "github.com/spf13/cobra"

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an existing task",
	Run:   func(cmd *cobra.Command, args []string) {},
}
