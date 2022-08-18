package cmd_test

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kakengloh/tsk/cmd"
	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/mock"
	"github.com/stretchr/testify/assert"
)

func Test_NewCommand(t *testing.T) {
	task := entity.Task{
		Name: "make coffee",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().CreateTask(task.Name, task.Priority, task.Status, "").Return(task, nil)

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
	task := entity.Task{
		Name:     "make coffee",
		Priority: entity.TaskPriorityMedium,
		Status:   entity.TaskStatusDoing,
		Comments: []string{"long black"},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().CreateTask(task.Name, task.Priority, task.Status, task.Comments[0]).Return(task, nil)

	buf := new(bytes.Buffer)

	newCmd := cmd.NewNewCommand(tr)
	newCmd.SetOut(buf)
	newCmd.SetErr(buf)
	newCmd.SetArgs([]string{
		"make coffee",
		"-p=medium",
		"-s=doing",
		"-c=long black",
	})

	err := newCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "1")
	assert.Contains(t, buf.String(), "make coffee")
	assert.Contains(t, buf.String(), "Medium")
	assert.Contains(t, buf.String(), "Doing")
}
