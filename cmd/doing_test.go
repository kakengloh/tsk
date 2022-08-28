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

func Test_DoingCommand(t *testing.T) {
	results := []repository.UpdateTaskStatusResult{
		{
			Task: entity.Task{
				Title: "make coffee",
			},
			FromStatus: entity.TaskStatusTodo,
			ToStatus:   entity.TaskStatusDoing,
		},
		{
			Task: entity.Task{
				Title: "feed my cat",
			},
			FromStatus: entity.TaskStatusDone,
			ToStatus:   entity.TaskStatusDoing,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock.NewMockTaskRepository(ctrl)
	tr.EXPECT().UpdateTaskStatus(entity.TaskStatusDoing, gomock.Any()).Return(results)

	buf := new(bytes.Buffer)

	todoCmd := cmd.NewDoingCommand(tr)
	todoCmd.SetOut(buf)
	todoCmd.SetErr(buf)
	todoCmd.SetArgs([]string{"1", "2"})

	err := todoCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, buf.String(), "make coffee: Todo -> Doing")
	assert.Contains(t, buf.String(), "feed my cat: Done -> Doing")
}
