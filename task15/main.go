package main

import (
	"fmt"
	"math/rand"
)

// That was rewritten:
// // 1. name the function according to its purpose of use
// // 2. Don't use global variable. If you want to change the value of variable in func body, use links and pointers
// // 3. Don't use the value like 1 << 10 = 1024 for making string about 100 symbols. Make the exact value

func main() {
	var justString string
	stringBuilder(&justString, 100)
	fmt.Println(justString)

}

func stringBuilder(justString *string, length int) {
	res := make([]rune, length)
	for i := 0; i < length; i++ {
		res[i] = rune('a' + rand.Int()%26)
	}
	*justString = string(res)
}
