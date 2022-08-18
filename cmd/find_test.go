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

func Test_FindCommand(t *testing.T) {
	tasks := entity.TaskList{
		entity.Task{
			ID:   1,
			Name: "make coffee",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().SearchTasks("coffee").Return(tasks, nil)

	buf := new(bytes.Buffer)

	findCmd := cmd.NewFindCommand(tr)
	findCmd.SetOut(buf)
	findCmd.SetErr(buf)
	findCmd.SetArgs([]string{"coffee"})

	err := findCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "1")
	assert.Contains(t, buf.String(), "make coffee")
}

func Test_FindCommandNoResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().SearchTasks("workout").Return(make([]entity.Task, 0), nil)

	buf := new(bytes.Buffer)

	findCmd := cmd.NewFindCommand(tr)
	findCmd.SetOut(buf)
	findCmd.SetErr(buf)
	findCmd.SetArgs([]string{"workout"})

	err := findCmd.Execute()
	assert.NoError(t, err)

	assert.Equal(t, "No task found with keyword \"workout\"\n", buf.String())
}
