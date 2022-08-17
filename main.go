package main

import (
	"log"

	"github.com/kakengloh/tsk/cmd"
	"github.com/kakengloh/tsk/driver"
	"github.com/kakengloh/tsk/repository"
)

func main() {
	db, err := driver.NewBolt()
	if err != nil {
		log.Fatalf("failed to connect to BoltDB: %s", err)
	}
	defer driver.CloseBolt()

	tr, err := repository.NewTaskRepository(db)
	if err != nil {
		log.Fatalf("failed to initialize task repository: %s", err)
	}

	cmd.Init(tr)
	cmd.Execute()
}
