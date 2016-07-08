package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	Log = func(msg string, d time.Duration) {
		if d != time.Second {
			t.Errorf("expected timeout to be one second, was %s", d)
		}
	}

	now = func() time.Time {
		return time.Date(2016, time.January, 12, 13, 14, 15, 0, time.UTC)
	}

	timer := New("test")

	now = func() time.Time {
		return time.Date(2016, time.January, 12, 13, 14, 16, 0, time.UTC)
	}

	timer.Stop()
}

func BenchmarkTimer(b *testing.B) {
	now = func() time.Time {
		return time.Now()
	}

	Log = func(msg string, d time.Duration) {

	}

	for i := 0; i < b.N; i++ {
		New("test").Stop()
	}
}
