package syncer

import (
	"errors"
	"log"
	"sync"
	"time"
)

var (
	ErrSyncTaskAlreadyExists = errors.New("a sync task with this name already exists")
	ErrTaskNotFound = errors.New("task not found")
)
type CallbackFunc func()

type task struct {
	stop chan bool
	ticker *time.Ticker
}

func (t *task) Stop() {
	t.stop <- true
	t.ticker.Stop()
}

type Syncer struct {
	tasks sync.Map
}

type Config struct {
	Delay time.Duration
}

func New() *Syncer {
	return &Syncer{}
}

func (m *Syncer) RegisterTask(name string, delay time.Duration, callback CallbackFunc) error {
	if _, ok := m.tasks.Load(name); ok {
		return ErrSyncTaskAlreadyExists
	}
	tsk := &task{
		ticker: time.NewTicker(delay),
		stop:  make(chan bool, 1),
	}
	m.tasks.Store(name, tsk)

	go func() {
		log.Printf("start syncing %s each %s", name, delay.String())
		callback()
		for {
			select {
			case <-tsk.stop:
				return
			case t := <-tsk.ticker.C:
				log.Printf("syncing %s at %s", name, t)
				callback()
			}
		}
	}()

	return nil
}

func (m *Syncer) StopTask(name string) error {
	if tsk, ok := m.tasks.Load(name); ok {
		t, ok := tsk.(*task)
		if !ok {
			return ErrTaskNotFound
		}

		t.Stop()
		m.tasks.Delete(name)
		return nil
	}
	return ErrTaskNotFound
}

func (m *Syncer) StopAllTasks() {
	m.tasks.Range(func(key, value interface{}) bool {
		if t, ok := value.(*task); ok {
			t.Stop()
		}
		m.tasks.Delete(key)
		return true
	})
}