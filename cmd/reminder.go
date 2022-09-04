package cmd

import (
	"fmt"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/services/reminder"
	"github.com/kakengloh/tsk/util"
	"github.com/spf13/cobra"
)

func NewReminderCommand(cr repository.ConfigRepository) *cobra.Command {
	reminderCmd := &cobra.Command{
		Use:   "reminder",
		Short: "Configure task reminder",
	}

	// tsk reminder start
	reminderCmd.AddCommand(newStartCommand())
	// tsk reminder stop
	reminderCmd.AddCommand(newStopCommand())
	// tsk reminder time
	reminderCmd.AddCommand(newTimeCommand(cr))

	return reminderCmd
}

func newStartCommand() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start reminder",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := reminder.Start()
			if err != nil {
				return fmt.Errorf("failed to start reminder: %w", err)
			}

			cmd.Println("Reminder started 🟢")

			return nil
		},
	}

	return startCmd
}

func newStopCommand() *cobra.Command {
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop reminder",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := reminder.Stop()
			if err != nil {
				return fmt.Errorf("failed to stop reminder: %w", err)
			}

			cmd.Println("Reminder stopped 🔴")

			return nil
		},
	}

	return stopCmd
}

func newTimeCommand(cr repository.ConfigRepository) *cobra.Command {
	timeCmd := &cobra.Command{
		Use:   "time",
		Short: "Reminder time before task due",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			durations, err := util.StringSliceToDurationSlice(args)
			if err != nil {
				return fmt.Errorf("failed to parse durations: %w", err)
			}

			err = cr.SetReminder(entity.ReminderConfig{
				Time: durations,
			})
			if err != nil {
				return fmt.Errorf("failed to update reminder time: %w", err)
			}

			cmd.Println("Reminder time updated ✅")

			return nil
		},
	}

	return timeCmd
}
