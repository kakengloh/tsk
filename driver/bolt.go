package driver

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func GetDataDir() (string, error) {
	dataroot := os.Getenv("XDG_DATA_HOME")

	if dataroot == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("$HOME directory not found: %w", err)
		}
		dataroot = filepath.Join(home, ".local", "share")
	}

	datadir := filepath.Join(dataroot, "tsk")
	err := os.MkdirAll(datadir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create tsk under %s directory: %w", dataroot, err)
	}

	return datadir, nil
}

func NewBolt() (*bbolt.DB, error) {
	datadir, err := GetDataDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(datadir, "bolt.db")

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

	return os.Remove(filepath.Join(home, ".tsk", "bolt.db"))
}
