package main

import (
	"fmt"
	"log"
)

// Tasks contains all avaliable task names.
var Tasks chan string

// NewTasks creates new tasks pool.
func NewTasks() chan string {
	return make(chan string, DefaultSizeOfATaskPool)
}

// Task contains usage statistics of a single task.
type Task struct {
	Name  string
	Usage uint64
}

func (t Task) String() string {
	return fmt.Sprintf("%s - %d", t.Name, t.Usage)
}

// GetNewTaskName returns new task name.
func GetNewTaskName() string {
	return RandomString(DefaultTaskNameLength, DefaultCharsList)
}

// GenerateNewTasks add new tasks of specified amount to a task pool.
func GenerateNewTasks(tasks chan string, n int) {
	for i := 0; i < n; i++ {
		task := GetNewTaskName()

		log.Printf("Seed task: '%s'\n", task)

		tasks <- task
	}
}

// EvictAndAddNewTask removes task from task pool and adds new task to tasks pool.
func EvictAndAddNewTask(tasks chan string) {
	var task string

	// evict only then tasks are avaliable
	if len(tasks) > 0 {
		task = <-tasks

		log.Printf("Evict task: '%s'\n", task)
	}

	task = GetNewTaskName()

	log.Printf("Add task: '%s'\n", task)

	tasks <- task
}
