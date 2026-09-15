package main

import "fmt"

func main() {
	var (
		age    int     = 19
		height float64 = 1.66
		name   string  = "Nastia"
		IsTrue bool    = true
	)

	fmt.Println("Привет! Меня зовут", name)
	fmt.Println("Мне", age)
	fmt.Println("Мой рост", height)
	fmt.Println("И это", IsTrue)
}
