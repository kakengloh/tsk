package cmd

import (
	"fmt"
	"strconv"

	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewRmCommand(tr repository.TaskRepository) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove an existing task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ids := []int{}

			for _, arg := range args {
				id, err := strconv.Atoi(arg)
				if err != nil {
					continue
				}
				ids = append(ids, id)
			}

			err := tr.DeleteTask(ids...)
			if err != nil {
				return fmt.Errorf("failed to delete tasks: %w", err)
			}

			fmt.Printf("\nTask(s) deleted ✅\n")

			return err
		},
	}

	return rmCmd
}
