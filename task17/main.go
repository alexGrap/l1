package main

import (
	"errors"
	"fmt"
	"sort"
)

// main
// // program is making binary search by the func

func main() {
	startArray := []int{8, 6, 9, 4, 12, 54, 36, 45, 78, 95, 32}
	target := 54
	sort.Ints(startArray)
	answer, err := binSearch(startArray, target)
	if err != nil {
		fmt.Println("Item not found")
		return
	}
	fmt.Printf("Index of %d is %d", target, answer)
}

// binary search func
// // function find the target in the sorted array by moving step-by-step from start and end to the middle of array

func binSearch(arr []int, target int) (int, error) {
	start := 0
	end := len(arr) - 1

	for start < end {
		median := (start + end) / 2
		if target == arr[median] {
			return median, nil
		}
		if arr[median] < target {
			start += 1
		} else {
			end -= 1
		}
	}
	return -1, errors.New("not found")
}
