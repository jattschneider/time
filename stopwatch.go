package time

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type task struct {
	Name  string
	Nanos time.Duration
}

//StopWatch allowing for timing of a number of tasks, exposing total running time and running time for each named task.
type StopWatch struct {
	tasks       []*task
	currentName string
	currentTime time.Time
	total       time.Duration
}

// Start a named task
func (sw *StopWatch) Start(taskName string) {
	sw.Stop()
	sw.currentName = fmt.Sprintf("[%s]", taskName)
	sw.currentTime = time.Now()
}

// Stop the current task
func (sw *StopWatch) Stop() {
	if sw.Running() {
		t := time.Now()
		elapsed := t.Sub(sw.currentTime)
		stopped := &task{Name: sw.currentName, Nanos: elapsed}
		sw.tasks = append(sw.tasks, stopped)
		sw.total = sw.total + elapsed
		sw.currentName = ""
	}
}

// Reset resets the StopWatch
func (sw *StopWatch) Reset() {
	sw.tasks = make([]*task, 0)
	sw.currentName = ""
	sw.currentTime = time.Time{}
	sw.total = 0
}

//Running determine whether this StopWatch is running
func (sw *StopWatch) Running() bool {
	return sw.currentName != ""
}

// Summary a short description of the total running time
func (sw *StopWatch) Summary() string {
	return fmt.Sprintf("StopWatch: running time = %v", sw.total)
}

func (sw *StopWatch) String() string {
	var sb strings.Builder
	sb.WriteString(sw.Summary())
	for _, t := range sw.tasks {
		sb.WriteString("; ")
		sb.WriteString(t.Name)
		sb.WriteString(" took ")
		sb.WriteString(t.Nanos.String())
		sb.WriteString(" = ")
		percent := math.Round(100.0 * float64(t.Nanos) / float64(sw.total))
		sb.WriteString(fmt.Sprintf("%v%%", percent))
	}
	return sb.String()
}
