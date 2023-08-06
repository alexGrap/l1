package main

import "fmt"

func main() {
	fmt.Printf("abcd: %t\n", isUnique("abcd"))
	fmt.Printf("abCdefA: %t\n", isUnique("abCdefA"))
	fmt.Printf("aabcd: %t\n", isUnique("aabcd"))
}

func isUnique(str string) bool {
	m := make(map[string]int)
	for _, el := range str {
		_, exist1 := m[string(el)]
		_, exist2 := m[string(el+32)]
		_, exist3 := m[string(el-32)]
		if !exist1 && !exist2 && !exist3 {
			m[string(el)] = 1
		} else {
			return false
		}
	}
	return true
}
