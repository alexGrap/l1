package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// main func is starting the publisher and reader goroutines and
// killing them after timeout from the start argument

func main() {
	//parse the argument
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(errors.New("uncorrectable program start. Write the time of program life like arg on starting "))
	}
	// making channel and waitGroup for publisher and reader routine
	chanel := make(chan int, 1)
	defer close(chanel)
	var wg sync.WaitGroup

	//making context with timeout for killing the goroutine after it
	ctx, timer := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer timer()

	//add the goroutine in waitGroup and starting them
	wg.Add(2)
	go writer(chanel, ctx, &wg)
	go reader(chanel, ctx, &wg)
	wg.Wait()

}

// WRITER
// // this func using like goRoutine and paste the growing value in the chanel with cooldown in 1 sec
func writer(chanel chan int, ctx context.Context, wg *sync.WaitGroup) {
	var i = 0

	for {

		select {
		case <-ctx.Done():
			log.Println("Writer is finished")
			wg.Done()
			return
		default:
			chanel <- i
			log.Printf("WRITER: Value %d was writted in chanel", i)
			time.Sleep(time.Second)
			i++

		}
	}
}

// READER
// // this func using like goRoutine and take the value from channel
func reader(chanel chan int, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			log.Println("reader is finished")
			wg.Done()
			return
		case value := <-chanel:
			log.Printf("READER: Value %d was read from chanel", value)
		}
	}
}

//in case <-ctx.Done(), which creating when timeout is over, above func are posting information
// in log and stops themselves
