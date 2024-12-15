package main

import (
	"fmt"
	diff "my_test/internal"
)

func main() {
	var a, b int
	fmt.Println("Введите числа для разности: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Printf("Результат разности чисел %d; %d равен %d", a, b, diff.Diff(a, b))
}
