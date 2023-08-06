package main

import "fmt"

// main
// // program reversing the input string
func main() {
	var input string
	fmt.Print("Input the string: ")
	fmt.Scan(&input)
	var result string
	for i := len(input) - 1; i >= 0; i-- {
		result += string(input[i])
	}
	fmt.Println(result)
}
