package scheduler_test

import (
	"context"
	"testing"
	"time"

	"github.com/joaopmatos/TaskUI/internal/scheduler"
)

var now = time.Now()

var baseTask = scheduler.Task{
	Path:              `\MyTasks\DailyBackup`,
	Name:              "DailyBackup",
	Status:            scheduler.TaskStateReady,
	NextRun:           now.Add(24 * time.Hour),
	HasLastRun:        true,
	LastRun:           now.Add(-24 * time.Hour),
	LastRunExitStatus: 0,
}

func TestListWithMockData(t *testing.T) {
	s := scheduler.NewFake(
		&baseTask,
		&scheduler.Task{
			Path:   `\Microsoft\Windows\Reboot`,
			Name:   "Reboot",
			Status: scheduler.TaskStateDisabled,
		},
		&scheduler.Task{
			Path:   `\MyTasks\HealthCheck`,
			Name:   "HealthCheck",
			Status: scheduler.TaskStateRunning,
		},
	)

	_, err := s.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestRunWithMockData(t *testing.T) {
	s := scheduler.NewFake(
		&baseTask,
	)

	err := s.Run(context.Background(), baseTask.Path)

	if err != nil {
		t.Fatal(err)
	}
}

func TestEnableWithMockData(t *testing.T) {
	s := scheduler.NewFake(
		&baseTask,
	)

	err := s.Enable(context.Background(), baseTask.Path)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDisableWithMockData(t *testing.T) {
	s := scheduler.NewFake(
		&baseTask,
	)

	err := s.Disable(context.Background(), baseTask.Path)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteWithMockData(t *testing.T) {
	s := scheduler.NewFake(
		&baseTask,
	)

	err := s.Disable(context.Background(), baseTask.Path)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCompleteWithMockData(t *testing.T) {
	runningTask := baseTask
	runningTask.Status = scheduler.TaskStateRunning

	s := scheduler.NewFake(
		&runningTask,
	)

	err := s.Run(context.Background(), baseTask.Path)

	if err != nil {
		t.Fatal(err)
	}
}
