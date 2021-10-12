package main

import (
	"fmt"
	"log"
	"os"
	"runner-app/runner"
	"time"
)

/*
Task => func(id int)
*/
func main() {

	timeout := 5 * time.Minute
	fmt.Println("PID = ", os.Getpid())
	r := runner.New(timeout)
	r.Add(createTask(1))
	r.Add(createTask(2))
	r.Add(createTask(3))
	r.Add(createTask(4))

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("terminating due to timeout")
			// timeout
		case runner.ErrInterrupt:
			log.Println("terminating due to interrupt")
			// interrupt
		}
	}

	log.Println("Process ended")
}

func createTask(id int) func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Minute)
	}
}
