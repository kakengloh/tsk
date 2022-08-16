package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kakengloh/tsk/entity"
	"github.com/teris-io/shortid"
	bolt "go.etcd.io/bbolt"
)

type TaskRepository struct {
	DB *bolt.DB
}

func NewTaskRepository(db *bolt.DB) (*TaskRepository, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Task"))
		return err
	})

	return &TaskRepository{db}, err
}

func (tr *TaskRepository) CreateTask(name string, priority entity.TaskPriority, status entity.TaskStatus, comment string) (entity.Task, error) {
	id, _ := shortid.Generate()

	comments := []entity.TaskComment{}
	if comment != "" {
		cid, _ := shortid.Generate()
		comments = append(comments, entity.TaskComment{
			ID:   cid,
			Text: comment,
		})
	}

	t := entity.Task{
		ID:        id,
		Name:      name,
		Priority:  priority,
		Status:    status,
		Comments:  comments,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := tr.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		b.Put([]byte(id), buf)

		return nil
	})

	if err != nil {
		return entity.Task{}, err
	}

	return t, nil
}

func (tr *TaskRepository) ListTasks() ([]entity.Task, error) {
	tasks := []entity.Task{}

	err := tr.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t entity.Task

			err := json.Unmarshal(v, &t)
			if err != nil {
				continue
			}

			tasks = append(tasks, t)
		}

		return nil
	})

	return tasks, err
}

func (tr *TaskRepository) GetTaskByID(id string) (entity.Task, error) {
	var t entity.Task

	err := tr.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		v := b.Get([]byte(id))
		if v == nil {
			return fmt.Errorf("task not found")
		}

		return json.Unmarshal(v, &t)
	})

	return t, err
}

func (tr *TaskRepository) DeleteTask(id string) error {
	err := tr.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		return b.Delete([]byte(id))
	})
	return err
}
