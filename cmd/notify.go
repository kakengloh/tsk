package cmd

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
	"github.com/xeonx/timeago"
	"golang.org/x/exp/slices"
)

func NewNotifyCommand(cr repository.ConfigRepository, tr repository.TaskRepository) *cobra.Command {
	notifyCmd := &cobra.Command{
		Use:   "notify",
		Short: "Notify on task due",
		RunE: func(cmd *cobra.Command, args []string) error {
			reminder, err := cr.GetReminder()
			if err != nil {
				return fmt.Errorf("failed to get reminders: %w", err)
			}

			var earliestReminderTime time.Duration
			remindInMinutes := make([]int, len(reminder.Time))

			for i, r := range reminder.Time {
				// Convert remind time to minute
				remindInMinutes[i] = int(r.Minutes())

				// Find earliest reminder time
				if r > earliestReminderTime {
					earliestReminderTime = r
				}
			}

			tasks, err := tr.ListTasksWithFilters(entity.TaskFilters{
				Due: earliestReminderTime,
			})
			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			for _, t := range tasks {
				min := int(time.Until(t.Due).Round(time.Minute).Minutes())
				if !slices.Contains(remindInMinutes, min) {
					continue
				}
				msg := timeago.English.Format(t.Due)
				beeep.Alert(t.Title, msg, "")
				cmd.Printf("Notified: %s\n", fmt.Sprintf("%s %s", t.Title, msg))
			}

			return nil
		},
	}

	return notifyCmd
}
