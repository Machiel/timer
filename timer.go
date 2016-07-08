// Package timer is a simple package which allows you to time tasks. For example
// database calls or other jobs.
package timer

import (
	"fmt"
	"time"
)

// Log outputs the result of the timer
var Log = func(msg string, d time.Duration) {
	fmt.Printf("%s took %s", msg, d)
}

var now = func() time.Time {
	return time.Now()
}

// New creates a new Timer
func New(msg string) *Timer {
	return &Timer{
		msg:   msg,
		start: now(),
	}
}

// Timer times a specific task
type Timer struct {
	start time.Time
	msg   string
}

// Stop ends the timer
func (t *Timer) Stop() {
	end := now()
	diff := end.Sub(t.start)
	Log(t.msg, diff)
}
