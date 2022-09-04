package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/util"
	"github.com/olekukonko/tablewriter"
	"github.com/xeonx/timeago"
)

type Printer struct {
	Stdout io.Writer
}

type OutputFormat string

const (
	OutputFormatTable = "table"
	OutputFormatJSON  = "json"
)

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
	created := timeago.English.Format(task.CreatedAt)
	// If duration is at least a week, show the exact date
	if time.Since(task.CreatedAt).Hours() >= 168 {
		created = task.CreatedAt.Format("2006-01-02")
	}

	// Calculate due
	due := ""
	if !task.Due.IsZero() {
		if time.Now().Before(task.Due) {
			// Before due
			due = timeago.English.Format(task.Due)
		} else {
			// Over due
			due = color.HiRedString(timeago.English.Format(task.Due))
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
		created := timeago.English.Format(t.CreatedAt)
		// If duration is at least a week, show the exact date
		if time.Since(t.CreatedAt).Hours() >= 168 {
			created = t.CreatedAt.Format("2006-01-02")
		}

		// Calculate due
		due := ""
		if !t.Due.IsZero() {
			if time.Now().Before(t.Due) {
				// Before due
				due = timeago.English.Format(t.Due)
			} else {
				// Over due
				due = color.HiRedString(timeago.English.Format(t.Due))
			}
		}

		table.Append([]string{strconv.Itoa(t.ID), t.Title, status, priority, created, due, notes})
	}

	fmt.Fprintln(p.Stdout)
	table.Render()
	fmt.Fprintln(p.Stdout)
}

func (p *Printer) PrintTaskListJSON(tasks []entity.Task) {
	results := make([]map[string]interface{}, len(tasks))

	for i, t := range tasks {
		// Format due
		due := ""
		if !t.Due.IsZero() {
			due = t.Due.Format("2006-01-02 15:04")
		}

		results[i] = map[string]interface{}{
			"id":       t.ID,
			"title":    t.Title,
			"priority": t.Priority.String(),
			"status":   t.Status.String(),
			"due":      due,
			"notes":    t.Notes,
			"created":  t.CreatedAt.Format("2006-01-02 15:04"),
		}
	}

	b, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("JSON marshal error")
		return
	}
	fmt.Fprintln(p.Stdout, string(b))
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
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
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
	summary += color.HiYellowString(fmt.Sprintf("%d doing / ", len(doing)))
	summary += color.HiGreenString(fmt.Sprintf("%d done", len(done)))
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
		s = color.HiYellowString(s)
	case entity.TaskStatusDone:
		s = color.HiGreenString(s)
	}

	return s
}

func ColoredPriority(status entity.TaskPriority) string {
	s := util.CapitalizeString(status.String())

	switch status {
	case entity.TaskPriorityLow:
		s = color.HiBlueString(s)
	case entity.TaskPriorityMedium:
		s = color.HiYellowString(s)
	case entity.TaskPriorityHigh:
		s = color.HiRedString(s)
	}

	return s
}
