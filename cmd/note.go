package cmd

import (
	"fmt"
	"strconv"

	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util/printer"
	"github.com/spf13/cobra"
)

func NewNoteCommand(tr repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "note",
		Short: "Add notes to an existing task",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			pt := printer.New(cmd.OutOrStdout())

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("task ID must be an integer: %w", err)
			}

			task, err := tr.AddNotes(id, args[1:]...)

			if err != nil {
				return fmt.Errorf("failed to add note(s): %w", err)
			}

			pt.PrintTask(task, "Note(s) added ✅")

			return nil
		},
	}
}
