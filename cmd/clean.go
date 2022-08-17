package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/driver"
	"github.com/spf13/cobra"
)

func NewCleanCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "Deletes all the data",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := driver.RemoveBolt()
			if err == nil {
				fmt.Println("Data deleted successfully")
			}
			return err
		},
	}
}
