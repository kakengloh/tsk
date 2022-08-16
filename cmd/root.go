package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tsk",
	Short: "tsk is a cli task management tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(mkCmd)
	rootCmd.AddCommand(edCmd)
	rootCmd.AddCommand(rmCmd)
}
