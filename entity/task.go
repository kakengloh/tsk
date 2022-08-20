package entity

import (
	"time"
)

type Task struct {
	ID        int
	Title     string
	Priority  TaskPriority
	Status    TaskStatus
	Due       time.Time
	Notes     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskPriority int

const (
	TaskPriorityNone TaskPriority = iota
	TaskPriorityLow
	TaskPriorityMedium
	TaskPriorityHigh
)

var TaskPriorityToString = map[TaskPriority]string{
	TaskPriorityNone:   "",
	TaskPriorityLow:    "low",
	TaskPriorityMedium: "medium",
	TaskPriorityHigh:   "high",
}

var TaskPriorityFromString = map[string]TaskPriority{
	"":       TaskPriorityNone,
	"low":    TaskPriorityLow,
	"medium": TaskPriorityMedium,
	"high":   TaskPriorityHigh,
}

func (p TaskPriority) String() string {
	return TaskPriorityToString[p]
}

type TaskStatus int

const (
	TaskStatusNone TaskStatus = iota
	TaskStatusTodo
	TaskStatusDoing
	TaskStatusDone
)

var TaskStatusToString = map[TaskStatus]string{
	TaskStatusNone:  "",
	TaskStatusTodo:  "todo",
	TaskStatusDoing: "doing",
	TaskStatusDone:  "done",
}

var TaskStatusFromString = map[string]TaskStatus{
	"":      TaskStatusNone,
	"todo":  TaskStatusTodo,
	"doing": TaskStatusDoing,
	"done":  TaskStatusDone,
}

func (s TaskStatus) String() string {
	return TaskStatusToString[s]
}

type TaskList []Task

func (tl TaskList) FilterByStatus(status TaskStatus) []Task {
	tasks := []Task{}

	for _, t := range tl {
		if t.Status == status {
			tasks = append(tasks, t)
		}
	}

	return tasks
}

func (tl TaskList) FilterByPriority(priority TaskPriority) []Task {
	tasks := []Task{}

	for _, t := range tl {
		if t.Priority == priority {
			tasks = append(tasks, t)
		}
	}

	return tasks
}
