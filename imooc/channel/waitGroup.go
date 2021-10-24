package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker2) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		//go func() { done <- true }()
		w.done()
	}

}

type worker2 struct {
	in   chan int
	done func()
}

func createWorker2(id int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo2() {
	var wg sync.WaitGroup
	wg.Add(20)

	var workers [10]worker2
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	//Channel as first-class citizen
	chanDemo2()

}
