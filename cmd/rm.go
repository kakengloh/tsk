package cmd

import (
	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewRmCmd(tr *repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "rm",
		Short: "Remove an existing task",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}
