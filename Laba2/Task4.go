package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Stroka(stroka string) int {
	return len([]rune(stroka))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку:) ")
	s, _ := reader.ReadString('\n')

	s = strings.TrimSpace(s)

	fmt.Println("Длина строки:", Stroka(s))
}
