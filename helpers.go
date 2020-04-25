package main

import (
	"math/rand"
	"strings"
	"time"
)

// DoEvery runs specified function every 'd' duration.
func DoEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

// RandomString generates random string of specified length and list of chars using default pseudo-random source.
func RandomString(length int, chars string) string {
	sb := strings.Builder{}
	sb.Grow(length)

	for i := 0; i < length; i++ {
		sb.WriteByte(
			chars[rand.Intn(len(chars))],
		)
	}

	return sb.String()
}
