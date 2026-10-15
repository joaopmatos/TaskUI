package main

import (
	"time"
)

// Types

// A Task's status according to TASK_STATE enumeration
type Status int

const (
	TaskStateUnknown Status = iota
	TaskStatePending
	TaskStateRunning
	TaskStateCompleted
	TaskStateFailed
)

func (s Status) String() string {
	states := [...]string{
		"UNKNOWN",
		"PENDING",
		"RUNNING",
		"COMPLETED",
		"FAILED",
	}

	// Make sure s exists in the array
	if s < TaskStateUnknown || s > TaskStateFailed {
		return "UNKNOWN"
	}
	return states[s]
}

type Task struct {
	Name              string    //The task's name
	LastRun           time.Time //When the last run was
	NextRun           time.Time //When the next run will be
	Status            Status    //Current status of the tasks
	LastRunExitStatus Status    //Exit status of the last run
}

func main() {
}
