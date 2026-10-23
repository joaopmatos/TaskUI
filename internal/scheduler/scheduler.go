package scheduler

import "time"

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
		"PENDING",
		"RUNNING",
		"COMPLETED",
		"FAILED",
	}

	// Make sure s exists in the array
	if s < TaskStateUnknown || s > TaskStateRunning {
		return "UNKNOWN"
	}
	return states[s]
}

type Task struct {
	ID                uint32            //The task's ID
	Name              string            //The task's name
	LastRun           time.Time         //When the last run was
	NextRun           time.Time         //When the next run will be
	Status            Status            //Current status of the tasks
	LastRunExitStatus LastRunExitStatus //Exit status of the last run
}

type LastRunExitStatus int32
