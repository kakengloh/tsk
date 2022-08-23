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

func Test_LsCommand(t *testing.T) {
	tasks := entity.TaskList{
		entity.Task{
			ID:       1,
			Title:    "make coffee",
			Status:   entity.TaskStatusTodo,
			Priority: entity.TaskPriorityLow,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().ListTasksWithFilters(gomock.Any()).Return(tasks, nil)

	buf := new(bytes.Buffer)

	lsCmd := cmd.NewLsCommand(tr)
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
	m.EXPECT().ListTasksWithFilters(gomock.Any()).Return(make(entity.TaskList, 0), nil)

	buf := new(bytes.Buffer)

	lsCmd := cmd.NewLsCommand(m)
	lsCmd.SetOut(buf)
	lsCmd.SetErr(buf)

	err := lsCmd.Execute()
	assert.NoError(t, err)

	expected := "No results found, try adjusting your filters to find what you're looking for!\n"
	assert.Equal(t, expected, buf.String())
}

func Test_LsCommandWithKeyword(t *testing.T) {
	tasks := entity.TaskList{
		entity.Task{
			ID:    1,
			Title: "make coffee",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().ListTasksWithFilters(gomock.Any()).Return(tasks, nil)

	buf := new(bytes.Buffer)

	lsCmd := cmd.NewLsCommand(tr)
	lsCmd.SetOut(buf)
	lsCmd.SetErr(buf)
	lsCmd.SetArgs([]string{"coffee"})

	err := lsCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "1")
	assert.Contains(t, buf.String(), "make coffee")
}
