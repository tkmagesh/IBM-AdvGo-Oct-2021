package main

import (
	"io"
	"log"
	"math/rand"
	"pooling-demo/pool"
	"sync"
	"sync/atomic"
	"time"
)

var maxGoroutines = 25
var pooledResources uint = 5

//Resource that has to be pooled
type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	println("Close: Connection", dbConn.ID)
	return nil
}

//counter to track the number of goroutines (clients)
var idCounter int32

//factory
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(maxGoroutines)

	pool, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Fatalln(err)
	}

	for idx := 0; idx < maxGoroutines; idx++ {
		go func(id int) {
			useResource(id, pool)
			wg.Done()
		}(idx)
	}
	wg.Wait()

	//verifying the "Acquire()" returns the resources from the pool
	wg.Add(3)
	for idx := 26; idx <= 28; idx++ {
		go func(id int) {
			useResource(idx, pool)
			wg.Done()
		}(idx)
	}
	wg.Wait()
}

func useResource(id int, pool *pool.Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	println("useResource:", id, "Connection", conn.(*dbConnection).ID)
}
