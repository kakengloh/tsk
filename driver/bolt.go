package driver

import (
	"time"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func NewBolt(path string) (*bolt.DB, error) {
	conn, err := bolt.Open(path, 0666, &bolt.Options{Timeout: time.Second})
	if err != nil {
		return nil, err
	}

	db = conn

	return db, nil
}

func CloseBolt() {
	db.Close()
}
