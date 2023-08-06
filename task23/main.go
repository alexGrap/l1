package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("Delete \"%d\" on 4 position from %v\n", arr[4], arr)
	arr = deleter(arr, 4)
	fmt.Printf("New arr: %v", arr)
}

func deleter(arr []int, pos int) []int {
	arr = append(arr[:pos], arr[pos+1:]...)
	return arr
}
