package printer

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/util"
	"github.com/olekukonko/tablewriter"
)

type Printer struct {
	Stdout io.Writer
}

func New(out io.Writer) *Printer {
	return &Printer{
		Stdout: out,
	}
}

func (p *Printer) PrintTask(task entity.Task, caption string) {
	table := tablewriter.NewWriter(p.Stdout)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	// Generate caption
	if caption != "" {
		table.SetCaption(true, caption)
	}

	// Generate headers
	table.SetAutoFormatHeaders(false)
	headers := []string{"ID", "Title", "Status", "Priority", "Created", "Due", "Notes"}
	headerColors := make([]tablewriter.Colors, len(headers))
	for i := range headers {
		headerColors[i] = tablewriter.Colors{tablewriter.Bold}
	}
	table.SetHeader(headers)
	table.SetHeaderColor(headerColors...)

	// Populate data
	// Status formatting
	status := ColoredStatus(task.Status)

	// Priority formatting
	priority := ColoredPriority(task.Priority)

	// Notes formatting
	notes := ""
	for i, c := range task.Notes {
		notes += fmt.Sprintf("%d) %s\n", i+1, c)
	}
	notes = strings.TrimSuffix(notes, "\n")

	// Calculate time ago
	// Show relative time by default
	created := time.Since(task.CreatedAt).Round(time.Second).String() + " ago"
	// If duration is at least a week, show the exact date
	if time.Since(task.CreatedAt).Hours() >= 168 {
		created = task.CreatedAt.Format("2006-01-02")
	}

	// Calculate due
	due := ""
	if !task.Due.IsZero() {
		duration := time.Until(task.Due).Round(time.Second)

		if duration.Milliseconds() > 0 {
			// Before due
			due = "in " + duration.String()

			if time.Until(task.Due).Hours() > 24 {
				due = task.Due.Format("2006-01-02 15:04")
			}
		} else {
			// Over due
			due = color.RedString("over " + duration.String()[1:])
		}
	}

	table.Append([]string{strconv.Itoa(task.ID), task.Title, status, priority, created, due, notes})

	fmt.Fprintln(p.Stdout)
	table.Render()
	fmt.Fprintln(p.Stdout)
}

func (p *Printer) PrintTaskList(tasks []entity.Task) {
	table := tablewriter.NewWriter(p.Stdout)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	// Generate caption
	caption := fmt.Sprintf("%d total", len(tasks))
	table.SetCaption(true, caption)

	// Generate headers
	headers := []string{"ID", "Title", "Status", "Priority", "Created", "Due", "Notes"}
	headerColors := make([]tablewriter.Colors, len(headers))
	for i := range headers {
		headerColors[i] = tablewriter.Colors{tablewriter.Bold}
	}
	table.SetAutoFormatHeaders(false)
	table.SetHeader(headers)
	table.SetHeaderColor(headerColors...)

	// Populate data
	for _, t := range tasks {
		// Status formatting
		status := ColoredStatus(t.Status)

		// Priority formatting
		priority := ColoredPriority(t.Priority)

		// Notes formatting
		notes := ""
		for i, n := range t.Notes {
			notes += fmt.Sprintf("%d) %s\n", i+1, n)
		}
		notes = strings.TrimSuffix(notes, "\n")

		// Calculate time ago
		// Show relative time by default
		created := time.Since(t.CreatedAt).Round(time.Second).String() + " ago"
		// If duration is at least a week, show the exact date
		if time.Since(t.CreatedAt).Hours() >= 168 {
			created = t.CreatedAt.Format("2006-01-02")
		}

		// Calculate due
		due := ""
		if !t.Due.IsZero() {
			duration := time.Until(t.Due).Round(time.Second)

			if duration.Milliseconds() > 0 {
				// Before due
				due = "in " + duration.String()

				if time.Until(t.Due).Hours() > 24 {
					due = t.Due.Format("2006-01-02 15:04")
				}
			} else {
				// Over due
				due = color.RedString("over " + duration.String()[1:])
			}
		}

		table.Append([]string{strconv.Itoa(t.ID), t.Title, status, priority, created, due, notes})
	}

	fmt.Fprintln(p.Stdout)
	table.Render()
	fmt.Fprintln(p.Stdout)
}

func (p *Printer) PrintTaskBoard(todo, doing, done entity.TaskList) {
	table := tablewriter.NewWriter(p.Stdout)
	table.SetAutoWrapText(false)

	// Generate headers
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{
		util.CapitalizeString(entity.TaskStatusTodo.String()),
		util.CapitalizeString(entity.TaskStatusDoing.String()),
		util.CapitalizeString(entity.TaskStatusDone.String()),
	})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)

	lists := []entity.TaskList{todo, doing, done}
	row := make([]string, len(lists))

	for i, tasks := range lists {
		content := ""
		for _, t := range tasks {
			content += fmt.Sprintf("%d) %s\n", t.ID, t.Title)
		}
		row[i] = strings.TrimSuffix(content, "\n")
	}

	table.Append(row)

	fmt.Fprintln(p.Stdout)
	table.Render()

	// Print summary
	summary := ""
	summary += color.HiBlueString(fmt.Sprintf("%d todo / ", len(todo)))
	summary += color.YellowString(fmt.Sprintf("%d doing / ", len(doing)))
	summary += color.GreenString(fmt.Sprintf("%d done", len(done)))
	fmt.Fprintln(p.Stdout, summary)
	fmt.Fprintln(p.Stdout)
}

func (p *Printer) PrintStatusUpdate(title string, from, to entity.TaskStatus, padding int) {
	fmt.Fprintf(p.Stdout, "\n%-*s: %s -> %s\n\n", padding, title, ColoredStatus(from), ColoredStatus(to))
}

func ColoredStatus(status entity.TaskStatus) string {
	s := util.CapitalizeString(status.String())

	switch status {
	case entity.TaskStatusTodo:
		s = color.HiBlueString(s)
	case entity.TaskStatusDoing:
		s = color.YellowString(s)
	case entity.TaskStatusDone:
		s = color.GreenString(s)
	}

	return s
}

func ColoredPriority(status entity.TaskPriority) string {
	s := util.CapitalizeString(status.String())

	switch status {
	case entity.TaskPriorityLow:
		s = color.HiBlueString(s)
	case entity.TaskPriorityMedium:
		s = color.YellowString(s)
	case entity.TaskPriorityHigh:
		s = color.RedString(s)
	}

	return s
}
