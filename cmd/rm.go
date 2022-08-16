package cmd

import (
	"fmt"
	"strconv"

	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewRmCmd(tr *repository.TaskRepository) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove an existing task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])

			if err != nil {
				return fmt.Errorf("task id must be an integer: %w", err)
			}

			err = tr.DeleteTask(id)
			if err != nil {
				return fmt.Errorf("failed to delete task: %w", err)
			}

			fmt.Printf("Task \"%d\" has been deleted\n", id)

			return nil
		},
	}

	return rmCmd
}
