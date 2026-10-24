package scheduler

import (
	"context"
	"time"
)

// Types

// A Task's status according to TASK_STATE enumeration
type Status int

const (
	TaskStateUnknown Status = iota
	TaskStateDisabled
	TaskStateQueued
	TaskStateReady
	TaskStateRunning
)

func (s Status) String() string {
	states := [...]string{
		"UNKNOWN",
		"DISABLED",
		"QUEUED",
		"READY",
		"RUNNING",
	}

	// Make sure s exists in the array
	if s < TaskStateUnknown || s > TaskStateRunning {
		return "UNKNOWN"
	}
	return states[s]
}

// Main structure: Task
type Task struct {
	Path              string            //The tasks's path: will be its identifier, e.g. \MyTasks\DailyBackup
	Name              string            //The task's name - only for display (leaf of the path)
	HasLastRun        bool              //Whether we have a last run or not
	LastRun           time.Time         //When the last run was
	NextRun           time.Time         //When the next run will be - zero value means "no next run".
	Status            Status            //Current status of the tasks
	LastRunExitStatus LastRunExitStatus //Exit status of the last run
}

type ListResult struct {
	Tasks    []Task
	Problems []TaskError
}

type TaskError struct {
	Path string
	Err  error
}

type lookupError struct {
	key string
	err error
}

func (err lookupError) Error() string {
	return err.key + ": " + err.err.Error()
}

type LastRunExitStatus uint32

type Scheduler interface {
	List(ctx context.Context) (result ListResult, err error)
	Run(ctx context.Context, name string) (err error)
	Enable(ctx context.Context, name string) (err error)
	Disable(ctx context.Context, name string) (err error)
	Delete(ctx context.Context, name string) (err error)
}
