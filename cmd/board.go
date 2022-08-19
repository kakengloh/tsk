package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util/printer"
	"github.com/spf13/cobra"
)

func NewBoardCommand(tr repository.TaskRepository) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "board",
		Short: "Display tasks in a Kanban board",
		RunE: func(cmd *cobra.Command, args []string) error {
			pt := printer.New(cmd.OutOrStdout())

			tasks, err := tr.ListTasks(entity.TaskStatusNone, entity.TaskPriorityNone, "")

			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			pt.PrintTaskBoard(
				tasks.FilterByStatus(entity.TaskStatusTodo),
				tasks.FilterByStatus(entity.TaskStatusDoing),
				tasks.FilterByStatus(entity.TaskStatusDone),
			)

			return nil
		},
	}

	return lsCmd
}
