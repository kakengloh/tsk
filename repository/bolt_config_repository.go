package repository

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/kakengloh/tsk/entity"
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

	err = cr.UpsertReminder(entity.ReminderConfig{
		Time: []time.Duration{15 * time.Minute},
	})

	return cr, err
}

func (cr *BoltConfigRepository) GetReminder() (entity.ReminderConfig, error) {
	var reminder entity.ReminderConfig

	err := cr.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Config"))
		buf := b.Get([]byte("reminder"))
		if buf == nil {
			return errors.New("reminder not set")
		}

		err := json.Unmarshal(buf, &reminder)

		return err
	})

	return reminder, err
}

func (cr *BoltConfigRepository) SetReminder(data entity.ReminderConfig) error {
	err := cr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Config"))
		v := b.Get([]byte("reminder"))
		if v == nil {
			return errors.New("reminder not set")
		}

		var r entity.ReminderConfig
		err := json.Unmarshal(v, &r)
		if err != nil {
			return err
		}

		if data.Time != nil {
			r.Time = data.Time
		}

		buf, err := json.Marshal(r)
		if err != nil {
			return err
		}

		b.Put([]byte("reminder"), buf)

		return err
	})

	return err
}

func (cr *BoltConfigRepository) UpsertReminder(data entity.ReminderConfig) error {
	err := cr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Config"))
		v := b.Get([]byte("reminder"))
		if v != nil {
			return nil
		}

		buf, err := json.Marshal(data)
		if err != nil {
			return err
		}

		b.Put([]byte("reminder"), buf)

		return err
	})

	return err
}
