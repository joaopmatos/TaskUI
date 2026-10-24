package scheduler

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrTaskNotFound = errors.New("task not found")

type FakeScheduler struct {
	mu    sync.RWMutex     // Read / Write manual Lock
	tasks map[string]*Task //Map of pointers to a Task
}

// Assigning *FakeScheduler(nil) to an anonymous variable allows us to force implementation of all methods
var _ Scheduler = (*FakeScheduler)(nil)

// Function takes either one or multiple tasks in the parameters
func NewFake(tasks ...*Task) *FakeScheduler {
	//Initialize the map (protects against nil)
	fake := &FakeScheduler{
		tasks: make(map[string]*Task),
	}
	for _, t := range tasks {
		if t == nil {
			continue
		}

		// store a copy: dereference and copy the struct
		copy := *t
		fake.tasks[copy.Path] = &copy
	}
	return fake
}

// Sorted by path
// error is fatal only
func (fake *FakeScheduler) List(ctx context.Context) (ListResult, error) {
	//Check if we have the context
	if err := ctx.Err(); err != nil {
		return ListResult{}, err
	}

	//Create a read lock
	fake.mu.RLock()

	//Defer the unlock - guarantees we unlock even if we return early (or panic).
	defer fake.mu.RUnlock()

	result := make([]Task, 0, len(fake.tasks))
	for _, task := range fake.tasks {
		result = append(result, *task)
	}

	return ListResult{Tasks: result}, nil
}

func (fake *FakeScheduler) Run(ctx context.Context, path string) error {
	//Check if we have the context
	if err := ctx.Err(); err != nil {
		return err
	}

	//Create a write lock
	fake.mu.Lock()

	//Defer the unlock - guarantees we unlock even if we return early (or panic).
	defer fake.mu.Unlock()

	task, ok := fake.tasks[path]
	if !ok {
		return ErrTaskNotFound
	}
	task.LastRun = time.Now()
	task.LastRunExitStatus = 0

	// For fake, keep status as-is or set to Ready after run
	if task.Status == TaskStateRunning {
		task.Status = TaskStateReady
	}
	return nil
}

func (fake *FakeScheduler) Enable(ctx context.Context, path string) error {
	//Check if we have the context
	if err := ctx.Err(); err != nil {
		return err
	}

	//Create a write lock
	fake.mu.Lock()

	//Defer the unlock - guarantees we unlock even if we return early (or panic).
	defer fake.mu.Unlock()

	task, ok := fake.tasks[path]
	if !ok {
		return ErrTaskNotFound
	}
	if task.Status == TaskStateDisabled {
		task.Status = TaskStateReady
	}
	return nil
}

func (fake *FakeScheduler) Disable(ctx context.Context, path string) error {
	//Check if we have the context
	if err := ctx.Err(); err != nil {
		return err
	}

	//Create a write lock
	fake.mu.Lock()

	//Defer the unlock - guarantees we unlock even if we return early (or panic).
	defer fake.mu.Unlock()

	task, ok := fake.tasks[path]
	if !ok {
		return ErrTaskNotFound
	}

	task.Status = TaskStateDisabled
	return nil
}

func (fake *FakeScheduler) Delete(ctx context.Context, path string) error {
	//Check if we have the context
	if err := ctx.Err(); err != nil {
		return err
	}

	//Create a write lock
	fake.mu.Lock()

	//Defer the unlock - guarantees we unlock even if we return early (or panic).
	defer fake.mu.Unlock()

	//_ is the blank identifier: we only care whether the key exists, not what it points to
	if _, ok := fake.tasks[path]; !ok {
		return ErrTaskNotFound
	}

	delete(fake.tasks, path)
	return nil
}
