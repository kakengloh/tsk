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
)

func PrintTasks(tasks []entity.Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	// Generate caption
	caption := fmt.Sprintf("%d in total", len(tasks))
	table.SetCaption(true, caption)

	// Generate headers
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"ID", "Name", "Status", "Priority", "Created", "Comments"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)

	// Populate data
	for _, t := range tasks {
		// Status formatting
		status := CapitalizeString(t.Status.String())
		switch t.Status {
		case entity.TaskStatusDoing:
			status = color.YellowString(status)
		case entity.TaskStatusDone:
			status = color.GreenString(status)
		}

		// Priority formatting
		priority := CapitalizeString(t.Priority.String())
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

func PrintTaskBoard(todo, doing, done entity.TaskList) {
	table := tablewriter.NewWriter(os.Stdout)

	// Generate caption
	caption := ""
	caption += fmt.Sprintf("%d todo / ", len(todo))
	caption += color.YellowString(fmt.Sprintf("%d doing / ", len(doing)))
	caption += color.GreenString(fmt.Sprintf("%d done", len(done)))
	table.SetCaption(true, caption)

	// Generate headers
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{
		CapitalizeString(entity.TaskStatusTodo.String()),
		CapitalizeString(entity.TaskStatusDoing.String()),
		CapitalizeString(entity.TaskStatusDone.String()),
	})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)

	// Calculate max rows
	max := len(todo)
	if len(doing) > max {
		max = len(doing)
	}
	if len(done) > max {
		max = len(done)
	}

	// Populate data
	for i := 0; i < max; i++ {
		row := []string{}

		if len(todo) > i {
			row = append(row, todo[i].Name)
		}
		if len(doing) > i {
			row = append(row, doing[i].Name)
		}
		if len(done) > i {
			row = append(row, done[i].Name)
		}

		table.Append(row)
	}

	table.Render()
}
