package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewRmCommand(tr repository.TaskRepository) *cobra.Command {
	rmCmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove an existing task",
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

			res := tr.BulkDeleteTasks(ids...)

			var err error

			success := []string{}

			for _, id := range ids {
				err = res[id]
				if err == nil {
					success = append(success, strconv.Itoa(id))
				} else {
					fmt.Printf("Failed to delete task \"%d\": %s\n", id, err)
				}
			}

			fmt.Printf("\nTask(s) [%s] deleted ❌\n", strings.Join(success, ", "))

			return err
		},
	}

	return rmCmd
}
