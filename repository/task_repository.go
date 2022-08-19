package repository

import "github.com/kakengloh/tsk/entity"

type TaskRepository interface {
	CreateTask(title string, priority entity.TaskPriority, status entity.TaskStatus, note string) (entity.Task, error)
	ListTasks() (entity.TaskList, error)
	SearchTasks(q string) (entity.TaskList, error)
	GetTaskByID(id int) (entity.Task, error)
	UpdateTask(id int, title string, priority entity.TaskPriority, status entity.TaskStatus) (entity.Task, error)
	UpdateTaskStatus(status entity.TaskStatus, ids ...int) []UpdateTaskStatusResult
	DeleteTask(id int) error
	BulkDeleteTasks(ids ...int) map[int]error
	AddNotes(id int, notes ...string) (entity.Task, error)
}

type UpdateTaskStatusResult struct {
	Task       entity.Task
	Err        error
	FromStatus entity.TaskStatus
	ToStatus   entity.TaskStatus
}
