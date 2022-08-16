package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/util"
	"github.com/teris-io/shortid"
	"go.etcd.io/bbolt"
)

type TaskRepository struct {
	DB *bbolt.DB
}

func NewTaskRepository(db *bbolt.DB) (*TaskRepository, error) {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Task"))
		return err
	})

	return &TaskRepository{db}, err
}

func (tr *TaskRepository) CreateTask(name string, priority entity.TaskPriority, status entity.TaskStatus, comment string) (entity.Task, error) {
	comments := []entity.TaskComment{}

	if comment != "" {
		cid, _ := shortid.Generate()
		comments = append(comments, entity.TaskComment{
			ID:   cid,
			Text: comment,
		})
	}

	var t entity.Task

	err := tr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))

		id, _ := b.NextSequence()

		t = entity.Task{
			ID:        int(id),
			Name:      name,
			Priority:  priority,
			Status:    status,
			Comments:  comments,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		b.Put(util.Itob(t.ID), buf)

		return nil
	})

	if err != nil {
		return entity.Task{}, err
	}

	return t, nil
}

func (tr *TaskRepository) ListTasks() ([]entity.Task, error) {
	tasks := []entity.Task{}

	err := tr.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		err := b.ForEach(func(k, v []byte) error {
			var t entity.Task

			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}

			tasks = append(tasks, t)
			return nil
		})

		return err
	})

	return tasks, err
}

func (tr *TaskRepository) GetTaskByID(id int) (entity.Task, error) {
	var t entity.Task

	err := tr.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		v := b.Get(util.Itob(id))
		if v == nil {
			return fmt.Errorf("task not found")
		}

		return json.Unmarshal(v, &t)
	})

	return t, err
}

func (tr *TaskRepository) DeleteTask(id int) error {
	err := tr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		return b.Delete(util.Itob(id))
	})
	return err
}
