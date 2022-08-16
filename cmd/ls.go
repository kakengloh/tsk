package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kakengloh/tsk/repository"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewLsCmd(tr *repository.TaskRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := tr.ListTasks()

			if len(tasks) == 0 {
				fmt.Println("You don't have any task yet, use the `tsk mk` command to make your first task!")
				return nil
			}

			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Status", "Priority", "Created"})

			for _, t := range tasks {
				status := cases.Title(language.English, cases.Compact).String(t.Status.String())
				priority := cases.Title(language.English, cases.Compact).String(t.Priority.String())
				since := time.Since(t.CreatedAt).Round(time.Second).String() + " ago"
				table.Append([]string{strconv.Itoa(t.ID), t.Name, status, priority, since})
			}

			table.Render()

			return nil
		},
	}
}
