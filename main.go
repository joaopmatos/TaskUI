package main

import (
	"os"
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
	ID                uint32    //The task's ID
	Name              string    //The task's name
	LastRun           time.Time //When the last run was
	NextRun           time.Time //When the next run will be
	Status            Status    //Current status of the tasks
	LastRunExitStatus Status    //Exit status of the last run
}

func main() {

	//Read the arguments
	arguments := os.Args

	//Check if we have any. If we don't have any arguments, we pass the --help output
	if len(arguments) <= 1 {
		init_tui()
		return
	}

	switch arguments[1] {
	case "list":
		list()
	case "run":
		run()
	case "help":
		help()
	default:
		init_tui()
	}

}

func init_tui() {
	println("TODO: implement TUI!")
}

func help() {
	const TextHelp = "TaskUI: lean Task Scheduler CLI tool.\n\nOptions\nlist: TaskUI list - shows a list of existing tasks\nrun: TaskUI run <id1> <id2> <id...> - runs the tasks enumerated in the parameters (separated by spaces)"

	println(TextHelp)
}

func list() {
	println("TODO: implement list function!")
}

func run() {
	println("TODO: implement list function!")
}
