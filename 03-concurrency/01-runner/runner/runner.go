package runner

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	timeout   <-chan time.Time
	tasks     []func(int)
	interrupt chan os.Signal
	complete  chan error
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func New(t time.Duration) *Runner {
	return &Runner{
		timeout:   time.After(t),
		interrupt: make(chan os.Signal),
		complete:  make(chan error),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for idx, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(idx)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		fmt.Println("interrupted and exiting")
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
