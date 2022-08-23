package repository

import (
	"bytes"
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

func (tr *BoltTaskRepository) CreateTask(title string, priority entity.TaskPriority, status entity.TaskStatus, due time.Time, note string) (entity.Task, error) {
	notes := []string{}
	if note != "" {
		notes = append(notes, note)
	}

	var t entity.Task

	err := tr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))

		id, _ := b.NextSequence()

		t = entity.Task{
			ID:        int(id),
			Title:     title,
			Priority:  priority,
			Status:    status,
			Due:       due,
			Notes:     notes,
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

func (tr *BoltTaskRepository) ListTasks(ids ...int) (entity.TaskList, error) {
	var tasks entity.TaskList

	err := tr.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		c := b.Cursor()

		var err error

		if len(ids) == 0 {
			// Retrieve all tasks
			err = b.ForEach(func(k, v []byte) error {
				var t entity.Task

				err = json.Unmarshal(v, &t)
				if err != nil {
					return err
				}

				tasks = append(tasks, t)
				return nil
			})
		} else {
			// Retrieve tasks with given IDs only
			containsID := func(ids []int, k []byte) bool {
				for _, id := range ids {
					if bytes.Equal(k, util.Itob(id)) {
						return true
					}
				}
				return false
			}

			for k, v := c.First(); k != nil && containsID(ids, k); k, v = c.Next() {
				var t entity.Task

				err = json.Unmarshal(v, &t)
				if err != nil {
					break
				}

				tasks = append(tasks, t)
			}
		}

		return err
	})

	return tasks, err
}

func (tr *BoltTaskRepository) ListTasksWithFilters(filters entity.TaskFilters) (entity.TaskList, error) {
	var tasks entity.TaskList

	err := tr.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		err := b.ForEach(func(k, v []byte) error {
			var t entity.Task

			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}

			// Filter by status
			if filters.Status != entity.TaskStatusNone {
				if filters.Status != t.Status {
					return nil
				}
			}

			// Filter by priority
			if filters.Priority != entity.TaskPriorityNone {
				if filters.Priority != t.Priority {
					return nil
				}
			}

			// Filter by keyword
			if filters.Keyword != "" {
				if !strings.Contains(strings.ToLower(t.Title), strings.ToLower(filters.Keyword)) {
					return nil
				}
			}

			// Filter by due
			if filters.Due.Seconds() > 0 {
				if t.Due.Second() == 0 {
					return nil
				}
				if time.Until(t.Due) >= filters.Due {
					return nil
				}
			}

			tasks = append(tasks, t)

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

func (tr *BoltTaskRepository) UpdateTask(id int, data entity.Task) (entity.Task, error) {
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

		if data.Title != "" {
			t.Title = data.Title
		}

		t.Priority = data.Priority
		t.Status = data.Status

		if !data.Due.IsZero() {
			t.Due = data.Due
		}

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

func (tr *BoltTaskRepository) DeleteTask(ids ...int) error {
	wg := &sync.WaitGroup{}
	ch := make(chan error, len(ids))

	for _, id := range ids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			err := tr.DB.Batch(func(tx *bbolt.Tx) error {
				b := tx.Bucket([]byte("Task"))
				return b.Delete(util.Itob(id))
			})

			ch <- err
		}(id)
	}

	wg.Wait()
	close(ch)

	var errMessages []string
	for msg := range ch {
		errMessages = append(errMessages, msg.Error())
	}

	var err error
	if len(errMessages) > 0 {
		err = fmt.Errorf(strings.Join(errMessages, "\n"))
	}

	return err
}

func (tr *BoltTaskRepository) AddNotes(id int, notes ...string) (entity.Task, error) {
	var t entity.Task

	err := tr.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Task"))

		v := b.Get(util.Itob(id))

		err := json.Unmarshal(v, &t)
		if err != nil {
			return err
		}

		t.Notes = append(t.Notes, notes...)

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		err = b.Put(util.Itob(id), buf)

		return err
	})

	return t, err
}
