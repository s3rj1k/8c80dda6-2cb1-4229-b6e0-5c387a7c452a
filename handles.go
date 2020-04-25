package main

import (
	"log"
	"net/http"
)

// RequestHandle returns next avaliable task name.
func RequestHandle(w http.ResponseWriter, r *http.Request) {
	// validate request method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	// THIS IS A HACK: trigger task generation then pool is empty
	if len(Tasks) == 0 {
		// this is needed to actually test concurrency using apache benchmark without hitting task starvation
		EvictAndAddNewTask(Tasks)
	}

	// get new task from pool
	task := <-Tasks

	// update task usage
	Stats.Inc(task)

	log.Printf("Issue task: '%s'\n", task)

	// write output
	if _, err := w.Write([]byte(task)); err != nil {
		panic(err)
	}

	// write newline
	if _, err := w.Write([]byte("\n")); err != nil {
		panic(err)
	}
}

// AdminRequestHandle returns next avaliable task name.
func AdminRequestHandle(w http.ResponseWriter, r *http.Request) {
	// validate request method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	// create output byte slice
	b := make([]byte, 0)

	// append usage statistics
	for _, el := range Stats.Usage() {
		b = append(b, []byte(el.String())...)
		b = append(b, []byte("\n")...)
	}

	// write output
	if _, err := w.Write(b); err != nil {
		panic(err)
	}
}
