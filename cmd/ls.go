package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util"
	"github.com/spf13/cobra"
)

func NewLsCommand(tr *repository.TaskRepository) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := tr.ListTasks()

			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			// Priority filter
			p, err := cmd.Flags().GetString("priority")
			if err != nil {
				return err
			}
			if p != "" {
				priority, ok := entity.TaskPriorityFromString[p]
				if !ok {
					return fmt.Errorf("invalid priority: %s, valid values are [low, medium, high]", p)
				}
				tasks = tasks.FilterByPriority(priority)
			}

			// Status filter
			s, err := cmd.Flags().GetString("status")
			if err != nil {
				return err
			}
			if s != "" {
				status, ok := entity.TaskStatusFromString[s]
				if !ok {
					return fmt.Errorf("invalid status: %s, valid values are [todo, doing, done]", s)
				}
				tasks = tasks.FilterByStatus(status)
			}

			if len(tasks) == 0 {
				fmt.Println("You don't have any task yet, use the `tsk new` command to create your first task!")
				return nil
			}

			util.PrintTasks(tasks)

			return nil
		},
	}

	lsCmd.PersistentFlags().StringP("priority", "p", "", "Filter by priority")
	lsCmd.PersistentFlags().StringP("status", "s", "", "Filter by status")

	return lsCmd
}
