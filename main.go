package main

import (
	"log"
	"time"

	"github.com/kakengloh/tsk/cmd"
	"github.com/kakengloh/tsk/repository"
	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("bolt.db", 0666, &bolt.Options{Timeout: time.Second})
	if err != nil {
		log.Fatalf("failed to connect to Bolt DB: %s", err)
	}
	defer db.Close()

	tr, err := repository.NewTaskRepository(db)
	if err != nil {
		log.Fatalf("failed to initialize task repository: %s", err)
	}

	cmd.Init(tr)
	cmd.Execute()
}
