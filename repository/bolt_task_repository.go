package repository

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/kakengloh/tsk/entity"
	"github.com/kakengloh/tsk/util"
	"go.etcd.io/bbolt"
)

type BoltTaskRepository struct {
	DB *bbolt.DB
}

func NewBoltTaskRepository(db *bbolt.DB) (*BoltTaskRepository, error) {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Task"))
		return err
	})

	return &BoltTaskRepository{db}, err
}

func (tr *BoltTaskRepository) CreateTask(name string, priority entity.TaskPriority, status entity.TaskStatus, comment string) (entity.Task, error) {
	comments := []string{}
	if comment != "" {
		comments = append(comments, comment)
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

func (tr *BoltTaskRepository) ListTasks() (entity.TaskList, error) {
	tasks := entity.TaskList{}

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

func (tr *BoltTaskRepository) SearchTasks(q string) (entity.TaskList, error) {
	tasks := entity.TaskList{}

	err := tr.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		err := b.ForEach(func(k, v []byte) error {
			var t entity.Task

			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}

			if strings.Contains(strings.ToLower(t.Name), strings.ToLower(q)) {
				tasks = append(tasks, t)
			}

			return nil
		})

		return err
	})

	return tasks, err
}

func (tr *BoltTaskRepository) GetTaskByID(id int) (entity.Task, error) {
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

func (tr *BoltTaskRepository) UpdateTask(id int, name string, priority entity.TaskPriority, status entity.TaskStatus) (entity.Task, error) {
	var t entity.Task

	err := tr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))

		v := b.Get(util.Itob(id))
		if v == nil {
			return fmt.Errorf("task not found")
		}

		err := json.Unmarshal(v, &t)
		if err != nil {
			return err
		}

		if name != "" {
			t.Name = name
		}

		t.Priority = priority
		t.Status = status

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		err = b.Put(util.Itob(id), buf)

		return err
	})

	return t, err
}

func (tr *BoltTaskRepository) UpdateTaskStatus(status entity.TaskStatus, ids ...int) []UpdateTaskStatusResult {
	wg := &sync.WaitGroup{}
	ch := make(chan UpdateTaskStatusResult, len(ids))

	for _, id := range ids {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			var r UpdateTaskStatusResult

			err := tr.DB.Batch(func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte("Task"))

				v := b.Get(util.Itob(id))
				if v == nil {
					return fmt.Errorf("task not found")
				}

				var t entity.Task

				err := json.Unmarshal(v, &t)
				if err != nil {
					return err
				}

				// Snapshot previous status
				fromStatus := t.Status

				// Assign new status
				t.Status = status

				buf, err := json.Marshal(t)
				if err != nil {
					return err
				}

				err = b.Put(util.Itob(id), buf)

				r = UpdateTaskStatusResult{
					Task:       t,
					FromStatus: fromStatus,
					ToStatus:   t.Status,
				}

				return err
			})

			r.Err = err

			ch <- r
		}(id)
	}

	wg.Wait()
	close(ch)

	res := []UpdateTaskStatusResult{}
	for r := range ch {
		res = append(res, r)
	}

	return res
}

func (tr *BoltTaskRepository) DeleteTask(id int) error {
	err := tr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		return b.Delete(util.Itob(id))
	})
	return err
}

func (tr *BoltTaskRepository) BulkDeleteTasks(ids ...int) map[int]error {
	type result struct {
		ID  int
		Err error
	}

	wg := &sync.WaitGroup{}
	ch := make(chan result, len(ids))

	for _, id := range ids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			err := tr.DB.Batch(func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte("Task"))
				return b.Delete(util.Itob(id))
			})

			ch <- result{id, err}
		}(id)
	}

	wg.Wait()
	close(ch)

	res := make(map[int]error)
	for msg := range ch {
		res[msg.ID] = msg.Err
	}

	return res
}

func (tr *BoltTaskRepository) AddComment(id int, comment string) (entity.Task, error) {
	var t entity.Task

	err := tr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))

		v := b.Get(util.Itob(id))

		err := json.Unmarshal(v, &t)
		if err != nil {
			return err
		}

		t.Comments = append(t.Comments, comment)

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		err = b.Put(util.Itob(id), buf)

		return err
	})

	return t, err
}