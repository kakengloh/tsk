package cmd

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

func NewNotifyCommand(cr repository.ConfigRepository, tr repository.TaskRepository) *cobra.Command {
	notifyCmd := &cobra.Command{
		Use:   "notify",
		Short: "Notify on task due",
		RunE: func(cmd *cobra.Command, args []string) error {
			reminders, err := cr.GetReminders()
			if err != nil {
				return fmt.Errorf("failed to get reminders: %w", err)
			}

			tasks, err := tr.ListTasksWithFilters(entity.TaskFilters{
				Due: time.Duration(10 * time.Minute),
			})
			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			for _, t := range tasks {
				min := int(time.Until(t.Due).Round(time.Minute).Minutes())
				if !slices.Contains(reminders, min) {
					continue
				}
				msg := fmt.Sprintf("in %d min", min)
				beeep.Alert(t.Title, msg, "")

				cmd.Printf("Notified: %s\n", fmt.Sprintf("%s %s", t.Title, msg))
			}

			return nil
		},
	}

	return notifyCmd
}
