package entity

import (
	"time"
)

type Task struct {
	ID        string
	Name      string
	Priority  TaskPriority
	Status    TaskStatus
	Comments  []TaskComment
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

type TaskComment struct {
	ID   string
	Text string
}
