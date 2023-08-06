package main

import (
	"fmt"
	"log"
	"sync"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

// main
// // program is creating 3 goRoutine, which increase value in struct and working parallels by mutex

func main() {
	counter := Counter{count: 0}

	//increase func
	operator := func(counter *Counter) {
		counter.mutex.Lock()
		counter.count++
		counter.mutex.Unlock()
	}

	//chan for waiting all routine done
	doneChan := make(chan int)

	go incrementation(&counter, operator, doneChan)
	go incrementation(&counter, operator, doneChan)
	go incrementation(&counter, operator, doneChan)
	doneWorkers := 0
	for {
		select {
		case <-doneChan:
			doneWorkers++
		}
		if doneWorkers == 3 {
			break
		}
	}
	fmt.Println(counter.count)
}

// incrementation func
// // func calls the small func to increase the struct value

func incrementation(counter *Counter, function func(counter *Counter), done chan int) {
	for i := 0; i < 10; i++ {
		function(counter)
	}
	log.Println("increment func done")
	done <- 1
}
