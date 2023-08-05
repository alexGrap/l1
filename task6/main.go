package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// FIRST WAY
// // close by the chanel

func firstWay(ch chan int) {
	for {
		select {
		case value := <-ch:
			log.Println(value)
		case <-ch:
			log.Println("FIRST WAY: routine is over")
			return
		}
	}
}

// SECOND WAY
// // using waitGroup
func secondWay() {
	log.Println("SECOND WAY: routine is starting")
	for {
	}
}

// THIRD WAY
// // by the context
func thirdWay(ctx context.Context, cancelChan chan bool) {
	log.Println("THIRD WAY: routine is starting")
	for {
		select {
		case <-ctx.Done():
			log.Println("THIRD WAY: routine is over")
			cancelChan <- true
			return
		}
	}
}

func main() {
	//using first way
	ch := make(chan int, 1)
	go firstWay(ch)
	time.Sleep(500 * time.Millisecond)
	ch <- 1
	close(ch)
	time.Sleep(1 * time.Second)

	//second way
	var wg sync.WaitGroup
	wg.Add(1)
	go secondWay()
	time.Sleep(3 * time.Second)
	wg.Done()
	log.Println("SECOND WAY: routine is over")

	//third way
	// using the ctx for killing routine and chanel for get info about routine
	canselChanel := make(chan bool, 1)
	ctx, cancel := context.WithCancel(context.Background())
	go thirdWay(ctx, canselChanel)
	time.Sleep(2 * time.Second)
	cancel()
	for {
		if <-canselChanel {
			break
		}
	}
}
