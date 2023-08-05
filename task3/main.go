package main

import (
	"fmt"
	"sync"
)

// "main" is making 5 goRoutine for calculating the squares sum of array values

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	chanel := make(chan int, 5)

	// operator func is posting in chanel result of calculation square value of current array value
	operator := func(wg *sync.WaitGroup, num int, chanel chan int) {
		chanel <- num * num
		wg.Done()
	}

	//loop making 5 goRoutine
	for _, el := range array {
		go operator(&wg, el, chanel)
		wg.Add(1)
	}
	wg.Wait()
	close(chanel)
	result := 0

	//loop is calculate the sum of squaring array value
	for el := range chanel {
		result += el
	}
	fmt.Printf("Result of calculation: %d", result)
}
