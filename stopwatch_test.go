package time

import (
	"fmt"
	"testing"
	"time"
)

func TestStopWatch(t *testing.T) {
	sw := &StopWatch{}

	t.Log(sw.Summary())
	t.Log(sw.String())

	for i := 0; i < 5; i++ {
		sw.Start(fmt.Sprintf("task #%v", i))
		time.Sleep(2 * time.Nanosecond)
		t.Logf("is running: %v", sw.Running())
	}
	sw.Stop()
	t.Log(sw.String())

	sw.Reset()
	for i := 0; i < 10; i++ {
		sw.Start(fmt.Sprintf("task #%v", i))
		sw.Stop()
	}
	t.Log(sw.String())
}
