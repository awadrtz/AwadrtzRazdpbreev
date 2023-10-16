package main

import (
	"fmt"
)

func main() {
	fmt.Println("Калькулятор")
	fmt.Println("Какое действие выполнить, (+, -, *, /, )")

	var action string
	fmt.Scan(&action)

	var a float64
	var b float64

	fmt.Println("Введите первое число")
	fmt.Scan(&a)

	fmt.Println("Введите второе число")
	fmt.Scan(&b)

	switch {
	case action == "+":
		fmt.Println("Сумма чисел = " + fmt.Sprint(a+b))

	case action == "-":
		fmt.Println("a-b = " + fmt.Sprint(a-b))

	case action == "*":
		fmt.Println("a * b = " + fmt.Sprint(a*b))

	case action == "/":
		fmt.Println("a / b = " + fmt.Sprint(a/b))

	}

}
