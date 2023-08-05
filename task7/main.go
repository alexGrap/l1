package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// main func is start 2 goRoutine for writing in map. The log has information about all
// steps and in the end print map for presentation result
// // program working 3 second and stop by context

func main() {
	//make context, map, closeChan for information of routine status and mutex for sync
	ctx, cansel := context.WithCancel(context.Background())
	m := make(map[int]string)
	closeChan := make(chan bool)
	var mutex sync.Mutex

	//starting 2 routine. First writing numbers from 1, second - from 100. Values in map - info about writer
	go writer(ctx, &mutex, m, closeChan, 1)
	go writer(ctx, &mutex, m, closeChan, 2)
	time.Sleep(3 * time.Second)
	cansel()
	doneCounter := 0

	//waiting closed status from routine
	for {
		if <-closeChan {
			doneCounter += 1
		}
		if doneCounter == 2 {
			log.Println("All writers done")
			break
		}
	}
	fmt.Println(m)
}

func writer(ctx context.Context, mutex *sync.Mutex, m map[int]string, ch chan bool, id int) {
	var (
		value int
		str   string
	)
	if id == 1 {
		str = "first writer"
		value = 1
	} else {
		str = "second writer"
		value = 100
	}
	for {
		select {
		case <-ctx.Done():
			ch <- true
			return
		default:
			mapWriter(mutex, m, value, str)
			log.Printf("%s is writed", str)
			value += 1
			time.Sleep(1 * time.Second)
		}
	}
}

func mapWriter(mutex *sync.Mutex, m map[int]string, value int, str string) {
	mutex.Lock()
	m[value] = str
	mutex.Unlock()
}
