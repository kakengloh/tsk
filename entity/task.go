package entity

import (
	"time"
)

type Task struct {
	ID        int
	Title     string
	Priority  TaskPriority
	Status    TaskStatus
	Notes     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskPriority int

const (
	TaskPriorityLow TaskPriority = iota
	TaskPriorityMedium
	TaskPriorityHigh
)

var TaskPriorityToString = map[TaskPriority]string{
	TaskPriorityLow:    "low",
	TaskPriorityMedium: "medium",
	TaskPriorityHigh:   "high",
}

var TaskPriorityFromString = map[string]TaskPriority{
	"low":    TaskPriorityLow,
	"medium": TaskPriorityMedium,
	"high":   TaskPriorityHigh,
}

func (p TaskPriority) String() string {
	return TaskPriorityToString[p]
}

type TaskStatus int

const (
	TaskStatusTodo TaskStatus = iota
	TaskStatusDoing
	TaskStatusDone
)

var TaskStatusToString = map[TaskStatus]string{
	TaskStatusTodo:  "todo",
	TaskStatusDoing: "doing",
	TaskStatusDone:  "done",
}

var TaskStatusFromString = map[string]TaskStatus{
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
