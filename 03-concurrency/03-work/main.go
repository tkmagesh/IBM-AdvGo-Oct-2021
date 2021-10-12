package main

import (
	"log"
	"time"
	"worker-demo/work"
)

var names = []string{
	"bob",
	"joe",
	"steve",
	"magesh",
	"suresh",
	"rajesh",
	"ramesh",
	"ganesh",
}

type namePrinter struct {
	name string
}

func (np *namePrinter) Task() {
	log.Println("Name Printer - Name : ", np.name)
	time.Sleep(2 * time.Second)
}

func main() {
	p := work.New(8)
	//var wg = sync.WaitGroup{}
	//wg.Add(10 * len(names))
	for i := 0; i < 10; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			p.Run(&np)
			//wg.Done()

		}
	}
	//wg.Wait()
	p.Shutdown()
}
