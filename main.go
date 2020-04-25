package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	// seed default pseudo-random source
	rand.Seed(time.Now().UnixNano())

	// debug logging: env DEBUG=TRUE ./tt
	if !strings.EqualFold(os.Getenv("DEBUG"), "TRUE") {
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	// initialize tasks channel
	Tasks = NewTasks()

	// initialize statistics object
	Stats = NewStatistics()

	// create initial tasks
	GenerateNewTasks(Tasks, DefaultNumberOfTasksOnStartup)

	// evict and add new task to tasks pool.
	go DoEvery(DefaultTimeForTaskRefreshAndEviction, func(_ time.Time) {
		EvictAndAddNewTask(Tasks)
	})

	http.HandleFunc("/request", RequestHandle)
	http.HandleFunc("/admin/requests", AdminRequestHandle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
