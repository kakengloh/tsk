package cmd

import (
	"fmt"
	"strconv"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewTodoCommand(tr *repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "todo",
		Short: "Mark task(s) as todo",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ids := []int{}

			for _, arg := range args {
				id, err := strconv.Atoi(arg)
				if err != nil {
					continue
				}
				ids = append(ids, id)
			}

			res := tr.UpdateTaskStatus(entity.TaskStatusTodo, ids...)

			var err error

			for k, v := range res {
				if v != nil {
					fmt.Printf("Failed to update task \"%d\": %s\n", k, v)
					if err == nil {
						err = fmt.Errorf("failed to update task: %w", v)
					}
				}
			}

			return err
		},
	}
}
