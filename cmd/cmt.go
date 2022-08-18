package cmd

import (
	"fmt"
	"strconv"

	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewCmtCommand(tr repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "cmt",
		Short: "Add comment to an existing task",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("task ID must be an integer: %w", err)
			}

			comment := args[1]

			_, err = tr.AddComment(id, comment)

			if err != nil {
				return fmt.Errorf("failed to add comment: %w", err)
			}

			fmt.Println("\nComment is added!")

			return nil
		},
	}
}
