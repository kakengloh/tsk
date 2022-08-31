package cmd

import (
	"github.com/kakengloh/tsk/repository"
	"github.com/spf13/cobra"
)

func NewConfigCommand(cr repository.ConfigRepository) *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Configure",
		RunE: func(cmd *cobra.Command, args []string) error {
			reminders, err := cmd.Flags().GetIntSlice("reminders")
			if err != nil {
				return err
			}

			err = cr.SetReminders(reminders)
			if err != nil {
				return err
			}

			cmd.Println("Config updated ✅")

			return nil
		},
	}

	configCmd.PersistentFlags().IntSlice("reminders", []int{}, "Set reminder(s) before task due (in minute)")

	return configCmd
}
