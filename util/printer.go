package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/kakengloh/tsk/entity"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PrintTasks(tasks []entity.Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{"ID", "Name", "Status", "Priority", "Created", "Comments"})

	for _, t := range tasks {
		// Status formatting
		status := cases.Title(language.English, cases.Compact).String(t.Status.String())
		switch t.Status {
		case entity.TaskStatusDoing:
			status = color.YellowString(status)
		case entity.TaskStatusDone:
			status = color.GreenString(status)
		}

		// Priority formatting
		priority := cases.Title(language.English, cases.Compact).String(t.Priority.String())
		switch t.Priority {
		case entity.TaskPriorityMedium:
			priority = color.YellowString(priority)
		case entity.TaskPriorityHigh:
			priority = color.RedString(priority)
		}

		// Comments formatting
		comments := ""
		if len(t.Comments) == 1 {
			comments = t.Comments[0].Text
		} else {
			for _, c := range t.Comments {
				comments += fmt.Sprintf("- %s\n", c.Text)
			}
			comments = strings.TrimSuffix(comments, "\n")
		}

		since := time.Since(t.CreatedAt).Round(time.Second).String() + " ago"
		table.Append([]string{strconv.Itoa(t.ID), t.Name, status, priority, since, comments})
	}

	table.Render()
}
