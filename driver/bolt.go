package driver

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func NewBolt() (*bbolt.DB, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("$HOME directory not found: %w", err)
	}

	err = os.MkdirAll(filepath.Join(home, ".tsk"), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create .tsk under $HOME directory: %w", err)
	}

	path := filepath.Join(home, ".tsk", "bolt.db")

	_, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

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

func RemoveBolt() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("$HOME directory not found: %w", err)
	}

	return os.RemoveAll(filepath.Join(home, ".tsk"))
}
