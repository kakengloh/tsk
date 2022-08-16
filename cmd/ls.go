package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewLsCmd(tr *repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := tr.ListTasks()

			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			for i, t := range tasks {
				fmt.Printf("%d. %s\n", i+1, t.Name)
			}

			return nil
		},
	}
}
