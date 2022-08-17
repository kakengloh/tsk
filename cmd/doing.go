package cmd

import (
	"fmt"
	"strconv"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util"
	"github.com/kakengloh/tsk/util/printer"
	"github.com/spf13/cobra"
)

func NewDoingCommand(tr repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "doing",
		Short: "Mark task(s) as doing",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pt := printer.New(cmd.OutOrStdout())

			ids := []int{}

			for _, arg := range args {
				id, err := strconv.Atoi(arg)
				if err != nil {
					continue
				}
				ids = append(ids, id)
			}

			res := tr.UpdateTaskStatus(entity.TaskStatusDoing, ids...)

			names := make([]string, len(res))
			for i, r := range res {
				names[i] = r.Task.Name
			}
			padding := util.MaxLen(names)

			for _, r := range res {
				if r.Err == nil {
					pt.PrintStatusUpdate(r.Task.Name, r.FromStatus, r.ToStatus, padding)
				} else {
					fmt.Printf("Failed to update task \"%s\": %s\n", r.Task.Name, r.Err)
				}
			}

			return nil
		},
	}
}
