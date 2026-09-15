package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите числа a, b, c: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println((a + b + c) / 3)

}
