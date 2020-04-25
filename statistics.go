package main

import (
	"sync"
)

// Statistics describes object that contains statuses for all tasks.
type Statistics struct {
	m  map[string]uint64
	mu *sync.RWMutex
}

// Stats contains statuses for all tasks.
var Stats *Statistics

// NewStatistics creates new Statistics object.
func NewStatistics() *Statistics {
	Statistics := new(Statistics)

	Statistics.m = make(map[string]uint64)
	Statistics.mu = new(sync.RWMutex)

	return Statistics
}

// Inc increments task count usage.
func (s *Statistics) Inc(task string) {
	s.mu.Lock()
	s.m[task]++
	s.mu.Unlock()
}

// Usage returns tasks usage statistics.
func (s *Statistics) Usage() []Task {
	usage := make([]Task, 0, len(s.m))

	for k, v := range s.m {
		usage = append(usage,
			Task{
				Name:  k,
				Usage: v,
			},
		)
	}

	return usage
}
