package cmd

import "github.com/spf13/cobra"

var mkCmd = &cobra.Command{
	Use:   "mk",
	Short: "Make a new task",
	Run:   func(cmd *cobra.Command, args []string) {},
}
