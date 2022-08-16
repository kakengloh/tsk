package cmd

import "github.com/spf13/cobra"

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}
