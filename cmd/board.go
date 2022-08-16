package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util"
	"github.com/spf13/cobra"
)

func NewBoardCmd(tr *repository.TaskRepository) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "board",
		Short: "Display tasks in a Kanban board",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := tr.ListTasks()

			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			util.PrintTaskBoard(
				tasks.FilterByStatus(entity.TaskStatusTodo),
				tasks.FilterByStatus(entity.TaskStatusDoing),
				tasks.FilterByStatus(entity.TaskStatusDone),
			)

			return nil
		},
	}

	return lsCmd
}
