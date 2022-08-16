package util

import (
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/kakengloh/tsk/entity"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PrintTasks(tasks []entity.Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Status", "Priority", "Created"})

	for _, t := range tasks {
		status := cases.Title(language.English, cases.Compact).String(t.Status.String())
		switch t.Status {
		case entity.TaskStatusDoing:
			status = color.YellowString(status)
		case entity.TaskStatusDone:
			status = color.GreenString(status)
		}

		priority := cases.Title(language.English, cases.Compact).String(t.Priority.String())
		switch t.Priority {
		case entity.TaskPriorityMedium:
			priority = color.YellowString(priority)
		case entity.TaskPriorityHigh:
			priority = color.RedString(priority)
		}

		since := time.Since(t.CreatedAt).Round(time.Second).String() + " ago"
		table.Append([]string{strconv.Itoa(t.ID), t.Name, status, priority, since})
	}

	table.Render()
}
