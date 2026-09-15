package main

import (
	"fmt"
	"math/rand/v2"
)

func Sum(a float64, b float64) float64 {
	return a + b

}

func Diff(a float64, b float64) float64 {
	return a - b
}

func main() {
	a := rand.Float64()
	fmt.Printf("Число a: %.2f \n", a)

	b := rand.Float64()
	fmt.Printf("Число b: %.2f \n", b)

	fmt.Printf("Сумма чисел с плавающей запятой: %.2f  \n", Sum(a, b))
	fmt.Printf("Разность чиселс плавающей запятой: %.2f \n", Diff(a, b))

}
