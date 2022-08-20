package cmd_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/kakengloh/tsk/cmd"
	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/mock"
	"github.com/stretchr/testify/assert"
)

func Test_NewCommand(t *testing.T) {
	task := entity.Task{
		Title:    "make coffee",
		Status:   entity.TaskStatusTodo,
		Priority: entity.TaskPriorityLow,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().CreateTask(task.Title, task.Priority, task.Status, gomock.Any(), "").Return(task, nil)

	buf := new(bytes.Buffer)

	newCmd := cmd.NewNewCommand(tr)
	newCmd.SetOut(buf)
	newCmd.SetErr(buf)
	newCmd.SetArgs([]string{"make coffee"})

	err := newCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "1")
	assert.Contains(t, buf.String(), "make coffee")
	assert.Contains(t, buf.String(), "Low")
	assert.Contains(t, buf.String(), "Todo")
}

func Test_NewCommandWithOptions(t *testing.T) {
	due, _ := time.ParseInLocation("2006-01-02 15:04", "2023-01-01 12:00", time.Local)
	task := entity.Task{
		Title:    "make coffee",
		Priority: entity.TaskPriorityMedium,
		Status:   entity.TaskStatusDoing,
		Notes:    []string{"long black"},
		Due:      due,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().CreateTask(task.Title, task.Priority, task.Status, due, task.Notes[0]).Return(task, nil)

	buf := new(bytes.Buffer)

	newCmd := cmd.NewNewCommand(tr)
	newCmd.SetOut(buf)
	newCmd.SetErr(buf)
	newCmd.SetArgs([]string{
		"make coffee",
		"-p=medium",
		"-s=doing",
		"-n=long black",
		"-d=2023-01-01 12:00",
	})

	err := newCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "1")
	assert.Contains(t, buf.String(), "make coffee")
	assert.Contains(t, buf.String(), "Medium")
	assert.Contains(t, buf.String(), "Doing")
	assert.Contains(t, buf.String(), "2023-01-01 12:00")
}
