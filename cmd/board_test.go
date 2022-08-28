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

func Test_BoardCommand(t *testing.T) {
	tasks := entity.TaskList{
		entity.Task{
			ID:       1,
			Title:    "make coffee",
			Status:   entity.TaskStatusTodo,
			Priority: entity.TaskPriorityLow,
		},
		entity.Task{
			ID:       2,
			Title:    "fix bug",
			Status:   entity.TaskStatusDoing,
			Priority: entity.TaskPriorityLow,
		},
		entity.Task{
			ID:       3,
			Title:    "play game",
			Status:   entity.TaskStatusDone,
			Priority: entity.TaskPriorityLow,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().ListTasks().Return(tasks, nil)

	buf := new(bytes.Buffer)

	boardCmd := cmd.NewBoardCommand(tr)
	boardCmd.SetOut(buf)
	boardCmd.SetErr(buf)

	err := boardCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "1) make coffee")
	assert.Contains(t, buf.String(), "2) fix bug")
	assert.Contains(t, buf.String(), "3) play game")

	assert.Contains(t, buf.String(), "1 todo / 1 doing / 1 done")
}
