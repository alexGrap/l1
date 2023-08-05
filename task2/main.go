package main

import (
	"fmt"
	"sync"
)

//main function make 5 goRoutine and working with them like with the local function

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	//local function, which printing the result of squaring the array value
	operator := func(wg *sync.WaitGroup, num int) {
		fmt.Println(num * num)
		wg.Done()
	}

	//the loop is starting goRoutine and adds them in the WaitGroup for waiting when it will be done
	for _, el := range array {
		go operator(&wg, el)
		wg.Add(1)

	}
	wg.Wait()
}
