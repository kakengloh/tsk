package repository

import (
	"errors"
	"time"

	"github.com/kakengloh/tsk/entity"
)

type TaskRepository interface {
	CreateTask(title string, priority entity.TaskPriority, status entity.TaskStatus, due time.Time, note string) (entity.Task, error)
	ListTasks(ids ...int) (entity.TaskList, error)
	ListTasksWithFilters(filters entity.TaskFilters) (entity.TaskList, error)
	GetTaskByID(id int) (entity.Task, error)
	UpdateTask(id int, data entity.Task) (entity.Task, error)
	UpdateTaskStatus(status entity.TaskStatus, ids ...int) []UpdateTaskStatusResult
	DeleteTask(id ...int) error
	AddNotes(id int, notes ...string) (entity.Task, error)
}

type UpdateTaskStatusResult struct {
	Task       entity.Task
	Err        error
	FromStatus entity.TaskStatus
	ToStatus   entity.TaskStatus
}

var ErrTaskNotFound = errors.New("task not found")
