package main

import (
	"strings"
	"testing"
)

func TestGetNewTaskName(t *testing.T) {
	name := GetNewTaskName()

	if len(name) != DefaultTaskNameLength {
		t.Errorf("GetNewTaskName() length = %d; want %d", len(name), DefaultTaskNameLength)
	}

	for pos, char := range name {
		if !strings.Contains(DefaultCharsList, string(char)) {
			t.Errorf("GetNewTaskName() returns unexpected character value at position %d, value = %c", pos, char)
		}
	}
}

func BenchmarkGetNewTaskName(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GetNewTaskName()
		}
	})
}
