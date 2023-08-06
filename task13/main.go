package main

import "fmt"

func main() {
	firstValue := 13
	secondValue := 25
	fmt.Printf("START POSITION\nFirst value : %d;   Second value: %d\n\n", firstValue, secondValue)
	firstValue, secondValue = secondValue, firstValue
	fmt.Printf("FIRST CALCULATION\nFirst value : %d;   Second value: %d\n\n", firstValue, secondValue)
	firstValue ^= secondValue
	secondValue = firstValue ^ secondValue
	firstValue = firstValue ^ secondValue
	fmt.Printf("SECOND CALCULATION\nFirst value : %d;   Second value: %d\n\n", firstValue, secondValue)
}
