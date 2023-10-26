package main

import (
	"fmt"
)

func main() {
	var firstNumber, secondNumber int
	var operation string

	fmt.Println("Введите первое число: ")
	firstNumber, err := InputNumber()
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	fmt.Println("Введите операцию (+, -, *, /): ")
	operation, err = InputOperation()
	if err != nil {
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /")
		return
	}

	fmt.Println("Введите второе число: ")
	secondNumber, err = InputNumber()
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	result, err := Evaluate(firstNumber, secondNumber, operation)
	if err != nil {
		fmt.Println("Ошибка: деление на ноль невозможно.")
		return
	}

	fmt.Println("Результат:", result)

}
