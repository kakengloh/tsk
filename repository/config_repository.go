package repository

import (
	"github.com/kakengloh/tsk/entity"
)

type ConfigRepository interface {
	GetReminder() (entity.ReminderConfig, error)
	SetReminder(data entity.ReminderConfig) error
}
