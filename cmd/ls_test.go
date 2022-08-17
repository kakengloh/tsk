package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kakengloh/tsk/cmd"
	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/mock"
	"github.com/stretchr/testify/assert"
)

var tasks = entity.TaskList{
	entity.Task{
		ID:       1,
		Name:     "make coffee",
		Priority: entity.TaskPriorityLow,
		Status:   entity.TaskStatusTodo,
	},
}

func Test_LsCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockTaskRepository(ctrl)
	m.EXPECT().ListTasks().Return(tasks, nil)

	buf := new(bytes.Buffer)

	lsCmd := cmd.NewLsCommand(m)
	lsCmd.SetOut(buf)
	lsCmd.SetErr(buf)

	err := lsCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "1")
	assert.Contains(t, buf.String(), "make coffee")
	assert.Contains(t, buf.String(), "Todo")
	assert.Contains(t, buf.String(), "Low")
}

func Test_LsCommandWithoutTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockTaskRepository(ctrl)
	m.EXPECT().ListTasks().Return(make(entity.TaskList, 0), nil)

	buf := new(bytes.Buffer)

	lsCmd := cmd.NewLsCommand(m)
	lsCmd.SetOut(buf)
	lsCmd.SetErr(buf)

	err := lsCmd.Execute()
	assert.NoError(t, err)

	expected := "You don't have any task yet, use the `tsk new` command to create your first task!"
	assert.Equal(t, expected, strings.TrimSpace(buf.String()))
}
