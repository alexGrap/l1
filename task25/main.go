package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start time: ", time.Now().Format("15:04:05"))
	Sleep(4 * time.Second)
	fmt.Println("End Time after 4 seconds sleeping: ", time.Now().Format("15:04:05"))
}

// Sleep
// // func making sleep on "time" seconds
func Sleep(duration time.Duration) {
	<-time.After(duration)
}
