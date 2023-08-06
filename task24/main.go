package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func Construction(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	first := Construction(5, 12)
	second := Construction(4, -14.5)
	fmt.Printf("The distance between points: %f", calculator(*first, *second))
}

func calculator(fir Point, sec Point) float64 {
	return math.Sqrt(math.Pow(fir.x-sec.x, 2) + math.Pow(fir.y-sec.y, 2))
}
