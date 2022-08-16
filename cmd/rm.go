package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewRmCmd(tr *repository.TaskRepository) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove an existing task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			err := tr.DeleteTask(id)
			if err != nil {
				return fmt.Errorf("failed to delete task: %w", err)
			}

			return nil
		},
	}

	return rmCmd
}
