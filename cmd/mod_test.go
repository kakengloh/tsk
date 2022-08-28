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

func Test_ModCommand(t *testing.T) {
	due, _ := time.ParseInLocation("2006-01-02 15:04", "2023-01-01 12:00", time.Local)

	task := entity.Task{
		ID:       1,
		Title:    "make coffee",
		Priority: entity.TaskPriorityMedium,
		Status:   entity.TaskStatusDoing,
		Notes:    []string{"long black"},
		Due:      due,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().UpdateTask(task.ID, gomock.Any()).Return(task, nil)

	buf := new(bytes.Buffer)

	modCmd := cmd.NewModCommand(tr)
	modCmd.SetOut(buf)
	modCmd.SetErr(buf)
	modCmd.SetArgs([]string{
		"1",
		"-p=medium",
		"-s=doing",
		"-d=2023-01-01 12:00",
	})

	err := modCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "1")
	assert.Contains(t, buf.String(), "make coffee")
	assert.Contains(t, buf.String(), "Medium")
	assert.Contains(t, buf.String(), "Doing")
	assert.Contains(t, buf.String(), "2023-01-01 12:00")
}
