package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util"
	"github.com/spf13/cobra"
)

func NewNewCommand(tr *repository.TaskRepository) *cobra.Command {
	mkCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Name
			name := args[0]

			// Priority
			p, err := cmd.Flags().GetString("priority")
			if err != nil {
				return err
			}
			priority, ok := entity.TaskPriorityFromString[p]
			if !ok {
				return fmt.Errorf("invalid priority: %s, valid values are [low, medium, high]", p)
			}

			// Status
			s, err := cmd.Flags().GetString("status")
			if err != nil {
				return err
			}
			status, ok := entity.TaskStatusFromString[s]
			if !ok {
				return fmt.Errorf("invalid status: %s, valid values are [todo, doing, done]", s)
			}

			// Comment
			c, err := cmd.Flags().GetString("comment")
			if err != nil {
				return err
			}

			// Create task
			t, err := tr.CreateTask(name, priority, status, c)

			if err != nil {
				return fmt.Errorf("failed to create task: %w", err)
			}

			util.PrintTask(t, "Task created ✅")

			return nil
		},
	}

	mkCmd.PersistentFlags().StringP("priority", "p", entity.TaskPriorityLow.String(), "Priority")
	mkCmd.PersistentFlags().StringP("status", "s", entity.TaskStatusTodo.String(), "Status")
	mkCmd.PersistentFlags().StringP("comment", "c", "", "Comment")

	return mkCmd
}
