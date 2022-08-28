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

func Test_DoneCommand(t *testing.T) {
	results := []repository.UpdateTaskStatusResult{
		{
			Task: entity.Task{
				Title: "make coffee",
			},
			FromStatus: entity.TaskStatusTodo,
			ToStatus:   entity.TaskStatusDone,
		},
		{
			Task: entity.Task{
				Title: "feed my cat",
			},
			FromStatus: entity.TaskStatusDoing,
			ToStatus:   entity.TaskStatusDone,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().UpdateTaskStatus(entity.TaskStatusDone, gomock.Any()).Return(results)

	buf := new(bytes.Buffer)

	todoCmd := cmd.NewDoneCommand(tr)
	todoCmd.SetOut(buf)
	todoCmd.SetErr(buf)
	todoCmd.SetArgs([]string{"1", "2"})

	err := todoCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "make coffee: Todo -> Done")
	assert.Contains(t, buf.String(), "feed my cat: Doing -> Done")
}
