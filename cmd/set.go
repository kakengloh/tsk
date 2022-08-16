package cmd

import (
	"fmt"
	"strconv"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewSetCmd(tr *repository.TaskRepository) *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set (update) an existing task",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("task ID must be an integer: %w", err)
			}

			t, err := tr.GetTaskByID(id)
			if err != nil {
				return fmt.Errorf("task not found")
			}

			// Name
			name := t.Name
			n, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			if n != "" {
				name = n
			}

			// Priority
			priority := t.Priority
			p, err := cmd.Flags().GetString("priority")
			if err != nil {
				return err
			}
			if p != "" {
				v, ok := entity.TaskPriorityFromString[p]
				if !ok {
					return fmt.Errorf("invalid priority: %s, valid values are [low, medium, high]", p)
				}
				priority = v
			}

			// Status
			status := t.Status
			s, err := cmd.Flags().GetString("status")
			if err != nil {
				return err
			}
			if s != "" {
				v, ok := entity.TaskStatusFromString[s]
				if !ok {
					return fmt.Errorf("invalid status: %s, valid values are [todo, doing, done]", s)
				}
				status = v
			}

			_, err = tr.UpdateTask(id, name, priority, status)

			return err
		},
	}

	setCmd.PersistentFlags().StringP("name", "n", "", "Name")
	setCmd.PersistentFlags().StringP("priority", "p", "", "Priority")
	setCmd.PersistentFlags().StringP("status", "s", "", "Status")

	return setCmd
}
