package main

import (
	"fmt"
	"strings"
)

// main
// // program is replacing words in the string from end to start

func main() {
	input := "snow dog sun"
	tmp := strings.Split(input, " ")
	fmt.Println(input)
	var result string
	for i := len(tmp) - 1; i >= 0; i-- {
		result += tmp[i]
		if i != 0 {
			result += " "
		}
	}
	fmt.Println(result)
}
