package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("abcd: %t\n", isUnique("abcd"))
	fmt.Printf("abCdefA: %t\n", isUnique("abCdefA"))
	fmt.Printf("aabcd: %t\n", isUnique("aabcd"))
}

func isUnique(str string) bool {
	m := make(map[string]int)
	str = strings.ToLower(str)
	for _, el := range str {
		_, exist1 := m[string(el)]
		if !exist1 {
			m[string(el)] = 1
		} else {
			return false
		}
	}
	return true
}
