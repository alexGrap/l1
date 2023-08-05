package main

import (
	"context"
	"fmt"
)

// main
// // func make 2 channels: 1st for writing, 2nd for reading and present result of calculating.
// // For calculating used goRoutine and function, which reading value and return result of
// // calculating in the 2nd channel

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	ctx, cansel := context.WithCancel(context.Background())
	go converter(ch1, ch2, ctx)
	for _, el := range arr {
		ch1 <- el
		fmt.Printf("Result for %d: %d\n", el, <-ch2)
	}
	cansel()
}

// CALCULATING FUNC

func converter(ch1 chan int, ch2 chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case value := <-ch1:
			ch2 <- value * 2
		}
	}
}
