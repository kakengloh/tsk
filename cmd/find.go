package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util"
	"github.com/spf13/cobra"
)

func NewFindCommand(tr *repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "find",
		Short: "Find tasks with keyword",
		RunE: func(cmd *cobra.Command, args []string) error {
			q := args[0]

			tasks, err := tr.SearchTasks(q)

			if len(tasks) == 0 {
				fmt.Printf("No tasks found with keyword \"%s\"\n", q)
				return nil
			}

			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			util.PrintTasks(tasks)

			return nil
		},
	}
}
