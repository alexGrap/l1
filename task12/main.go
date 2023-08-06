package main

import "fmt"

// main
// // program is creating the start array, and by the func making the set
func main() {
	startArray := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(stringSet(startArray))
}

// stringSer
// // func is making the map and by it making the unique array of startArray values

func stringSet(arr []string) []string {
	var result []string
	m := make(map[string]int)
	for _, el := range arr {
		m[el] = 1
	}
	for key, _ := range m {
		result = append(result, key)
	}
	return result
}
