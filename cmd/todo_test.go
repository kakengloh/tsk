package cmd_test

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kakengloh/tsk/cmd"
	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/mock"
	"github.com/kakengloh/tsk/repository"
	"github.com/stretchr/testify/assert"
)

func Test_TodoCommand(t *testing.T) {
	results := []repository.UpdateTaskStatusResult{
		{
			Task: entity.Task{
				Title: "make coffee",
			},
			FromStatus: entity.TaskStatusDoing,
			ToStatus:   entity.TaskStatusTodo,
		},
		{
			Task: entity.Task{
				Title: "feed my cat",
			},
			FromStatus: entity.TaskStatusDone,
			ToStatus:   entity.TaskStatusTodo,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().UpdateTaskStatus(entity.TaskStatusTodo, gomock.Any()).Return(results)

	buf := new(bytes.Buffer)

	todoCmd := cmd.NewTodoCommand(tr)
	todoCmd.SetOut(buf)
	todoCmd.SetErr(buf)
	todoCmd.SetArgs([]string{"1", "2"})

	err := todoCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "make coffee: Doing -> Todo")
	assert.Contains(t, buf.String(), "feed my cat: Done -> Todo")
}
