package repository

type ConfigRepository interface {
	GetReminders() ([]int, error)
	SetReminders(minutes []int) error
}
