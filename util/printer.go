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

func ColoredStatus(status entity.TaskStatus) string {
	s := CapitalizeString(status.String())

	switch status {
	case entity.TaskStatusTodo:
		s = color.BlackString(s)
	case entity.TaskStatusDoing:
		s = color.YellowString(s)
	case entity.TaskStatusDone:
		s = color.GreenString(s)
	}

	return s
}

func ColoredPriority(status entity.TaskPriority) string {
	s := CapitalizeString(status.String())

	switch status {
	case entity.TaskPriorityLow:
		s = color.BlackString(s)
	case entity.TaskPriorityMedium:
		s = color.YellowString(s)
	case entity.TaskPriorityHigh:
		s = color.GreenString(s)
	}

	return s
}

func PrintTask(task entity.Task, caption string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	// Generate caption
	if caption != "" {
		table.SetCaption(true, caption)
	}

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
	// Status formatting
	status := ColoredStatus(task.Status)

	// Priority formatting
	priority := ColoredPriority(task.Priority)

	// Comments formatting
	comments := ""
	for i, c := range task.Comments {
		comments += fmt.Sprintf("%d) %s\n", i+1, c)
	}
	comments = strings.TrimSuffix(comments, "\n")

	// Calculate time ago
	// Show relative time by default
	since := time.Since(task.CreatedAt).Round(time.Second).String() + " ago"
	// If duration is at least a week, show the exact date
	if time.Since(task.CreatedAt).Hours() >= 168 {
		since = task.CreatedAt.Format("02/01/2006")
	}

	table.Append([]string{strconv.Itoa(task.ID), task.Name, status, priority, since, comments})

	table.Render()
}

func PrintTaskList(tasks []entity.Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	// Generate caption
	caption := fmt.Sprintf("%d total", len(tasks))
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
		status := ColoredStatus(t.Status)

		// Priority formatting
		priority := ColoredPriority(t.Priority)

		// Comments formatting
		comments := ""
		for i, c := range t.Comments {
			comments += fmt.Sprintf("%d) %s\n", i+1, c)
		}
		comments = strings.TrimSuffix(comments, "\n")

		// Calculate time ago
		// Show relative time by default
		since := time.Since(t.CreatedAt).Round(time.Second).String() + " ago"
		// If duration is at least a week, show the exact date
		if time.Since(t.CreatedAt).Hours() >= 168 {
			since = t.CreatedAt.Format("02/01/2006")
		}

		table.Append([]string{strconv.Itoa(t.ID), t.Name, status, priority, since, comments})
	}

	table.Render()
}

func PrintTaskBoard(todo, doing, done entity.TaskList) {
	table := tablewriter.NewWriter(os.Stdout)

	// Generate headers
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{
		CapitalizeString(entity.TaskStatusTodo.String()),
		CapitalizeString(entity.TaskStatusDoing.String()),
		CapitalizeString(entity.TaskStatusDone.String()),
	})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
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
			row = append(row, fmt.Sprintf("%d) %s", todo[i].ID, todo[i].Name))
		} else {
			row = append(row, "")
		}

		if len(doing) > i {
			row = append(row, fmt.Sprintf("%d) %s", doing[i].ID, doing[i].Name))
		} else {
			row = append(row, "")
		}

		if len(done) > i {
			row = append(row, fmt.Sprintf("%d) %s", done[i].ID, done[i].Name))
		} else {
			row = append(row, "")
		}

		table.Append(row)
	}

	table.Render()

	// Print summary
	summary := ""
	summary += color.BlackString(fmt.Sprintf("%d todo / ", len(todo)))
	summary += color.YellowString(fmt.Sprintf("%d doing / ", len(doing)))
	summary += color.GreenString(fmt.Sprintf("%d done", len(done)))
	fmt.Println(summary)
}

func PrintStatusUpdate(name string, from, to entity.TaskStatus, padding int) {
	fmt.Printf("%-*s: %s -> %s\n", padding, name, ColoredStatus(from), ColoredStatus(to))
}
