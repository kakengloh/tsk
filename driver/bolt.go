package driver

import (
	"time"

	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func NewBolt(path string) (*bbolt.DB, error) {
	conn, err := bbolt.Open(path, 0666, &bbolt.Options{Timeout: time.Second})
	if err != nil {
		return nil, err
	}

	db = conn

	return db, nil
}

func CloseBolt() {
	db.Close()
}
