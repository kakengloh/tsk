package main

import (
	"log"

	"github.com/kakengloh/tsk/cmd"
	"github.com/kakengloh/tsk/driver"
	"github.com/kakengloh/tsk/repository"
)

func main() {
	// Database
	db, err := driver.NewBolt()
	if err != nil {
		log.Fatalf("failed to connect to BoltDB: %s", err)
	}
	defer driver.CloseBolt()

	// Config repository
	cr, err := repository.NewBoltConfigRepository(db)
	if err != nil {
		log.Fatalf("failed to initialize config repository: %s", err)
	}

	// Task repository
	tr, err := repository.NewBoltTaskRepository(db)
	if err != nil {
		log.Fatalf("failed to initialize task repository: %s", err)
	}

	cmd.Init(cr, tr)
	cmd.Execute()
}
