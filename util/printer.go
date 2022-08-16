package util

import (
	"os"
	"strconv"
	"time"

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
		priority := cases.Title(language.English, cases.Compact).String(t.Priority.String())
		since := time.Since(t.CreatedAt).Round(time.Second).String() + " ago"
		table.Append([]string{strconv.Itoa(t.ID), t.Name, status, priority, since})
	}

	table.Render()
}
