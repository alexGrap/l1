package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

//// main function reads an argument from the initial parameters and uses them to create N goRoutines that
//// will be killed by the user command Ctrl+C.

func main() {
	// parsing the arguments
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(errors.New("uncorrectable program start. Write the count of routine like arg on starting "))
	}
	if arg < 1 {
		arg = 1
	}

	//making the main and closing channels
	chanel := make(chan int, arg)
	closedChanel := make(chan bool, arg)
	var wg sync.WaitGroup

	// creating the context for using it like command for stop the goRoutines
	ctx, ctxCancel := context.WithCancel(context.Background())
	wg.Add(arg)
	for i := 0; i < arg; i++ {
		go worker(chanel, i, ctx, closedChanel)
	}

	// creating the anonymous func to paste the random value in the main chanel
	go func() {
		for {
			chanel <- rand.Int()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// creating the chanel which working with systems command
	shotDown := make(chan os.Signal, 1)
	signal.Notify(shotDown, syscall.SIGINT, syscall.SIGTERM)
	<-shotDown
	ctxCancel()
	countOfClosed := 0

	//creating the "while" loop to collect messages about finishing all goRoutines
	for {
		value := <-closedChanel

		if value {
			countOfClosed += 1
		}
		if countOfClosed == arg {
			break
		}
	}
	close(chanel)

	log.Println("Program end")

}

// WORKER func
// // This func is used in goRoutine and in the loop is reading values for presentation.
// // When user press Ctrl+C the ctx call the cansel func, and worker send information in the closingChanel,
// // that it's finish them work, and breaking the loop
func worker(chanel <-chan int, num int, ctx context.Context, closed chan bool) {
	log.Printf("Worker №%d is working", num)
	for {
		select {
		case value := <-chanel:
			log.Printf("Worker №%d: %d\n", num, value)
		case <-ctx.Done():
			closed <- true
			log.Printf("Worker №%d is finish", num)
			return
		}
	}
}
