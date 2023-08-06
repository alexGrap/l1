package main

import "fmt"

// main
// // program is sorting the start array by quicksort algorithm

func main() {
	arr := []int{1, 6, 9, 80, 17, 65, 79, 105, 32, 15, 74, 7}
	fmt.Println(sort(arr, 0, len(arr)-1))
}

// recursion func to present the result of sort

func sort(arr []int, left int, right int) []int {
	if left < right {
		var tmp int
		arr, tmp = tmpFunc(arr, left, right)
		arr = sort(arr, left, tmp-1)
		arr = sort(arr, tmp+1, right)
	}
	return arr
}

// sort function
// is going into array and replace the values

func tmpFunc(arr []int, left, right int) ([]int, int) {
	tmp := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < tmp {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}
