package cmd

import (
	"fmt"
	"time"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/repository"
	"github.com/kakengloh/tsk/util/printer"
	"github.com/spf13/cobra"
)

func NewLsCommand(tr repository.TaskRepository) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List tasks",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pt := printer.New(cmd.OutOrStdout())

			// Status filter
			var status entity.TaskStatus = 0
			s, err := cmd.Flags().GetString("status")
			if err != nil {
				return err
			}
			if s != "" {
				val, ok := entity.TaskStatusFromString[s]
				if !ok {
					return fmt.Errorf("invalid status: %s, valid values are [todo, doing, done]", s)
				}
				status = val
			}

			// Priority filter
			var priority entity.TaskPriority = 0
			p, err := cmd.Flags().GetString("priority")
			if err != nil {
				return err
			}
			if p != "" {
				val, ok := entity.TaskPriorityFromString[p]
				if !ok {
					return fmt.Errorf("invalid priority: %s, valid values are [low, medium, high]", p)
				}
				priority = val
			}

			// Keyword filter
			keyword := ""
			if len(args) > 0 {
				keyword = args[0]
			}

			// Due filter
			due, err := cmd.Flags().GetDuration("due")
			if err != nil {
				return nil
			}

			// Output format
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}

			tasks, err := tr.ListTasksWithFilters(entity.TaskFilters{
				Status:   status,
				Priority: priority,
				Keyword:  keyword,
				Due:      due,
			})

			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			if len(tasks) == 0 {
				cmd.Println("No results found, try adjusting your filters to find what you're looking for!")
				return nil
			}

			switch format {
			case printer.OutputFormatJSON:
				pt.PrintTaskListJSON(tasks)
			case printer.OutputFormatTable:
				pt.PrintTaskList(tasks)
			default:
				return fmt.Errorf("invalid output format: %s, valid values are [table, json]", format)
			}

			return nil
		},
	}

	lsCmd.PersistentFlags().StringP("status", "s", "", "Filter by status (todo / doing / done)")
	lsCmd.PersistentFlags().StringP("priority", "p", "", "Filter by priority (low / medium / high)")
	lsCmd.PersistentFlags().DurationP("due", "d", time.Duration(0), "Filter by due (eg: 30m / 1h / 2d)")
	lsCmd.PersistentFlags().StringP("format", "f", printer.OutputFormatTable, "Output format (table / json)")

	return lsCmd
}
