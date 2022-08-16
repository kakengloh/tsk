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

			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"", "ID", "Name", "Status", "Created"})

			for i, t := range tasks {
				index := strconv.Itoa(i + 1)
				status := cases.Title(language.English, cases.Compact).String(t.Status.String())
				since := time.Since(t.CreatedAt).Round(time.Second).String() + " ago"
				table.Append([]string{index, t.ID, t.Name, status, since})
			}

			table.Render()

			return nil
		},
	}
}
