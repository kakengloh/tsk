package cmd

import (
	"fmt"
	"strconv"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util"
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

			names := make([]string, len(res))
			for i, r := range res {
				names[i] = r.Task.Name
			}
			padding := util.MaxLen(names)

			for _, r := range res {
				if r.Err == nil {
					util.PrintStatusUpdate(r.Task.Name, r.FromStatus, r.ToStatus, padding)
				} else {
					fmt.Printf("Failed to update task \"%s\": %s\n", r.Task.Name, r.Err)
				}
			}

			return nil
		},
	}
}
