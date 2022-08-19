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

func NewDoneCommand(tr repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "done",
		Short: "Mark task(s) as done",
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

			res := tr.UpdateTaskStatus(entity.TaskStatusDone, ids...)

			titles := make([]string, len(res))
			for i, r := range res {
				titles[i] = r.Task.Title
			}
			padding := util.MaxLen(titles)

			for _, r := range res {
				if r.Err == nil {
					pt.PrintStatusUpdate(r.Task.Title, r.FromStatus, r.ToStatus, padding)
				} else {
					fmt.Printf("Failed to update task \"%s\": %s\n", r.Task.Title, r.Err)
				}
			}

			return nil
		},
	}
}
