package entity

import "time"

type Config struct {
	Reminder ReminderConfig
}

type ReminderConfig struct {
	Time []time.Duration
}
