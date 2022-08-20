package cmd

import (
	"fmt"
	"time"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util/printer"
	"github.com/spf13/cobra"
)

func NewNewCommand(tr repository.TaskRepository) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pt := printer.New(cmd.OutOrStdout())

			// Title
			title := args[0]

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

			// Note
			n, err := cmd.Flags().GetString("note")
			if err != nil {
				return err
			}

			// Due
			d, err := cmd.Flags().GetString("due")
			if err != nil {
				return err
			}
			var due time.Time
			if d != "" {
				duration, err := time.ParseDuration(d)
				if err == nil {
					due = time.Now().Add(duration)
				} else {
					val, err := time.ParseInLocation("2006-01-02 15:04", d, time.Local)
					if err != nil {
						return err
					}
					due = val
				}
			}

			// Create task
			t, err := tr.CreateTask(title, priority, status, due, n)

			if err != nil {
				return fmt.Errorf("failed to create task: %w", err)
			}

			pt.PrintTask(t, "Task created ✅")

			return nil
		},
	}

	newCmd.PersistentFlags().StringP("priority", "p", entity.TaskPriorityLow.String(), "Priority")
	newCmd.PersistentFlags().StringP("status", "s", entity.TaskStatusTodo.String(), "Status")
	newCmd.PersistentFlags().StringP("note", "n", "", "Note")
	newCmd.PersistentFlags().StringP("due", "d", "", "Due")

	return newCmd
}
