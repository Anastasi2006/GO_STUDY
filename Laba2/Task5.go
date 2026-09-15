package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func printRectangle(S Rectangle) int {
	return S.Width * S.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	S := printRectangle(rect)
	fmt.Println("Площадь прямоугольника:", S)
}
