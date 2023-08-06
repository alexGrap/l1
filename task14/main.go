package main

import "fmt"

// main
// // program is detecting the type of variable by the func

func main() {
	fmt.Println(typeDetector(int64(85)))
	fmt.Println(typeDetector(float32(56)))
	fmt.Println(typeDetector(byte(87)))
	fmt.Println(typeDetector(true))
}

// typeDetector
// // this func return the string witch contains the name of variable type.
// // for detecting is using the string flag %T witch print the type of variable
func typeDetector(x interface{}) string {
	return fmt.Sprintf("%T", x)
}
