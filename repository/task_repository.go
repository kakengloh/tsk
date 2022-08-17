package repository

import "github.com/kakengloh/tsk/entity"

type TaskRepository interface {
	CreateTask(name string, priority entity.TaskPriority, status entity.TaskStatus, comment string) (entity.Task, error)
	ListTasks() (entity.TaskList, error)
	SearchTasks(q string) (entity.TaskList, error)
	GetTaskByID(id int) (entity.Task, error)
	UpdateTask(id int, name string, priority entity.TaskPriority, status entity.TaskStatus) (entity.Task, error)
	UpdateTaskStatus(status entity.TaskStatus, ids ...int) []UpdateTaskStatusResult
	DeleteTask(id int) error
	BulkDeleteTasks(ids ...int) map[int]error
	AddComment(id int, comment string) (entity.Task, error)
}

type UpdateTaskStatusResult struct {
	Task       entity.Task
	Err        error
	FromStatus entity.TaskStatus
	ToStatus   entity.TaskStatus
}
