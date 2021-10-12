package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	mutex     sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrNegativePoolSize = errors.New("Size value too small.")
var ErrPoolClosed = errors.New("Pool has been closed.")

//Creating an instance of the pool
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, ErrNegativePoolSize
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case resource, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		return resource, nil
	default:
		log.Println("Acquire: New Resource")
		return p.factory()
	}
}

func (p *Pool) Release(resource io.Closer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		resource.Close()
		return
	}
	select {
	case p.resources <- resource:
		log.Println("Release: ", "Resource in queue")

	default:
		//if the pool is already full, then we close the resource
		log.Println("Release: ", "Pool is full, closing resource")
		resource.Close()
	}
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for resource := range p.resources {
		resource.Close()
	}
}
