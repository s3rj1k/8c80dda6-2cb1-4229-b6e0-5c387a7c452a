package main

import (
	"math"
	"time"
)

const (
	// DefaultCharsList defines default list of chars for task name.
	DefaultCharsList = "abcdefghijklmnopqrstuvwxyz"

	// DefaultTaskNameLength defines default length of task name.
	DefaultTaskNameLength = 2

	// DefaultNumberOfTasksOnStartup defines how many new tasks are generated before service is started.
	DefaultNumberOfTasksOnStartup = 50

	// DefaultTimeForTaskRefreshAndEviction defines amount of wait time for runnig task eviction on tasks pool.
	DefaultTimeForTaskRefreshAndEviction = 200 * time.Millisecond

	// DefaultSizeOfATaskPool maximum size of tasks pool.
	DefaultSizeOfATaskPool = math.MaxInt16
)
