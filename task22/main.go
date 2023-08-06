package main

import (
	"fmt"
	"math/big"
)

func main() {
	first := big.NewInt(1 << 21)
	second := big.NewInt(1 << 22)
	res := big.NewInt(0)
	fmt.Println("Add operation: ", res.Add(first, second))
	res = big.NewInt(0)
	fmt.Println("Sub operation: ", res.Sub(first, second))
	res = big.NewInt(0)
	fmt.Println("Mul operation: ", res.Mul(first, second))
	res = big.NewInt(0)
	fmt.Println("Div operation: ", res.Div(second, first))

}
