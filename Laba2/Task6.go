package main

import (
	"fmt"
	"math/rand/v2"
)

func average2(a int, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	a := rand.IntN(100)
	fmt.Println("Первое число:", a)
	b := rand.IntN(100)
	fmt.Println("Второе число:", b)
	fmt.Println("Среднее двух чисел:", average2(a, b))
}
