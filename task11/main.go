package main

import "fmt"

// main
// // this program is creating 2 arrays, and by the func "findCrossing" making the new one with intersection of thirst two

func main() {
	firstArr := []float64{25.6, 32.7, 45.9, 12.65, -9.97, 20.03, 19.89}
	secondArr := []float64{25.6, 45.9, 13.58, 98.63, 20.03, 16.45, 19.89}
	fmt.Println(findCrossing(firstArr, secondArr))
}

// findCrossing
// // this func is making the map for collect the unique values and score of unUnique, and after that making the new
// // one result array

func findCrossing(first []float64, second []float64) []float64 {
	m := make(map[float64]int)
	var result []float64
	for _, el := range first {
		m[el] = 1
	}
	for _, el := range second {
		if _, isExist := m[el]; isExist {
			m[el] += 1
		}
	}
	for key, value := range m {
		if value > 1 {
			result = append(result, key)
		}
	}
	return result
}
