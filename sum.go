package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(sub(2, 3))
	fmt.Println(times(2, 3))
	fmt.Println(divide(2, 3))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func divide(a int, b int) float64 {
	result := float64(a) / float64(b)
	ratio := math.Pow(10, float64(1))
	return math.Floor(result*ratio) / ratio
}
