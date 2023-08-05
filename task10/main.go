package main

import "fmt"

//main
// // func is working with example array.
// // in loop func is making variable "round" equals to tens degrees. After that is checking on exist in map,
// // and after that appending value in array or creating new one note

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.9, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)
	for _, el := range arr {
		round := int(el/10) * 10
		if round == 0 {
			if el < 0 {
				round = -1
			} else {
				round = 1
			}
		}
		if _, isExist := m[round]; isExist {
			m[round] = append(m[round], el)
		} else {
			m[round] = make([]float64, 0)
			m[round] = append(m[round], el)
		}
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
