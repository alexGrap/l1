package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

// main
// // program takes position and value from argument and stand value into position-bit

func main() {
	var variable int64
	variable = 23
	fmt.Printf("Start value: %d\n", variable)
	position, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(errors.New("uncorrected start program"))
	}
	bit, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(errors.New("uncorrected start program"))
	}
	if bit == 1 {
		variable |= int64(bit) << int64(position)
	} else if bit == 0 {
		variable &= ^int64(bit) << int64(position)
	} else {
		log.Fatal(errors.New("uncorrected start parameters"))
	}
	fmt.Printf("Result value: %d", variable)
}
