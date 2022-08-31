package repository

import (
	"encoding/json"

	"go.etcd.io/bbolt"
)

type BoltConfigRepository struct {
	DB *bbolt.DB
}

func NewBoltConfigRepository(db *bbolt.DB) (*BoltConfigRepository, error) {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Config"))
		return err
	})
	if err != nil {
		return nil, err
	}

	cr := &BoltConfigRepository{db}

	err = cr.SetReminders([]int{15})
	return cr, err
}

func (cr *BoltConfigRepository) GetReminders() ([]int, error) {
	var minutes []int

	err := cr.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Config"))
		buf := b.Get([]byte("reminders"))
		if buf == nil {
			// Default to 15 minute reminder if not configured
			minutes = []int{15}
			return nil
		}
		return json.Unmarshal(buf, &minutes)
	})

	return minutes, err
}

func (cr *BoltConfigRepository) SetReminders(minutes []int) error {
	err := cr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Config"))
		buf, err := json.Marshal(minutes)
		if err != nil {
			return err
		}
		err = b.Put([]byte("reminders"), buf)
		return err
	})

	return err
}
