package cmd

import (
	"fmt"
	"strings"

	"github.com/kakengloh/tsk/driver"
	"github.com/spf13/cobra"
)

func NewCleanCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "Clear all data",
		RunE: func(cmd *cobra.Command, args []string) error {
			var input string

			cmd.Println("Are you sure you want to clear all data? This action cannot be undone (y/n)")
			cmd.Print(">> ")
			fmt.Scanf("%s", &input)

			if strings.ToLower(input) != "y" {
				return nil
			}

			err := driver.RemoveBolt()
			if err == nil {
				fmt.Println("Data cleared successfully ✅")
			}
			return err
		},
	}
}
