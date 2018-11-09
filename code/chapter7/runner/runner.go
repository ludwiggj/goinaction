// Package runner manages the running and lifetime of a process.
// Example provided with help from Gabriel Aszalos.
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of tasks within a given timeout and can be
// shut down on an operating system interrupt.
type Runner struct {
	// interrupt channel reports a signal from the operating system.
	interrupt chan os.Signal

	// A Signal represents an operating system signal. The usual underlying implementation
	// is operating system-dependent: on Unix it is syscall.Signal.

	// type Signal interface {
	//   String() string
	//   Signal() // to distinguish from other Stringers
	// }

	// complete channel reports that processing is done.
	// If an error occurs, it’s reported back via an error interface value sent through the channel.
	// If no error occurs, the value of nil is sent as the error interface value.
	complete chan error

	// timeout reports that time has run out.
	// This is a "receive-only" channel
	timeout <-chan time.Time

	// tasks holds a set of functions that are executed synchronously in index order.
	tasks []func(int)
}

// ErrTimeout is returned when a value is received on the timeout.
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt is returned when an event from the OS is received.
var ErrInterrupt = errors.New("received interrupt")

// New returns a new ready-to-use Runner.
func New(d time.Duration) *Runner {
	return &Runner{
		// buffered channel of length 1
		interrupt: make(chan os.Signal, 1),

		// unbuffered channel
		complete: make(chan error),

		// The After function returns a channel of type time.Time
		// The runtime will send a time.Time value on this channel after the specified duration has elapsed.
		timeout: time.After(d),
	}
}

// Add attaches tasks to the Runner. A task is a function that takes an int ID.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events.
func (r *Runner) Start() error {
	// We want to receive all interrupt based signals.

	signal.Notify(r.interrupt, os.Interrupt)
	// Run the different tasks on a different goroutine.
	go func() {
		r.complete <- r.run()
	}()

	// Select (https://tour.golang.org/concurrency/5)
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case.
	// It chooses one at random if multiple are ready.
	select {
	// Signalled when processing is done - blocking operation
	case err := <-r.complete:
		return err
	// Signalled when we run out of time - presumably a blocking operation
	case <-r.timeout:
		return ErrTimeout
	}
}

// run executes each registered task.
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// Check for an interrupt signal from the OS.
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// Execute the registered task.
		task(id)
	}
	return nil
}

// gotInterrupt verifies if the interrupt signal has been issued.
func (r *Runner) gotInterrupt() bool {

	// Non-blocking receive, as per https://gobyexample.com/non-blocking-channel-operations
	select {
	// Signaled when an interrupt event is sent.
	case <-r.interrupt:
		// Stop receiving any further signals.
		signal.Stop(r.interrupt)
		return true
	// Continue running as normal.
	default:
		return false
	}
}
