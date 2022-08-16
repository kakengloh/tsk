package cmd

import (
	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewEdCmd(tr *repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "ed",
		Short: "Edit an existing task",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}
