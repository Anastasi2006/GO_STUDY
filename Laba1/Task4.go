package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	a := rand.IntN(100)
	fmt.Println("Число а:", a)
	b := rand.IntN(90)
	fmt.Println("Число b:", b)

	fmt.Println("Сумма чисел:", a+b)
	fmt.Println("Разность чисел:", a-b)
	fmt.Println("Результат деления:", a/b)
	fmt.Println("Результат умножения", a*b)

}
