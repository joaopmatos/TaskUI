package main

import (
	"time"
)

type Task struct {
	name    string
	lastRun time.Time
	nextRun time.Time
	status  uint8
}

func main() {
}
